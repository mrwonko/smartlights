package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/mrwonko/smartlights/config"
	"github.com/mrwonko/smartlights/internal/protocol"
	"github.com/simulatedsimian/joystick"

	rpio "github.com/stianeikeland/go-rpio/v4"
)

const (
	red   = rpio.Pin(13)
	green = rpio.Pin(19)
	blue  = rpio.Pin(18)
)

func main() {
	const piEnv = "PI_ID"
	// TODO: parse from human-readable string?
	pi, err := strconv.Atoi(os.Getenv(piEnv))
	if err != nil {
		log.Fatalf("failed to parse $%s: %s", piEnv, err)
	}
	err = rpio.Open()
	if err != nil {
		log.Fatalf("failed to init go-rpio: %s", err)
	}
	defer rpio.Close()

	red.Output()
	green.Output()
	blue.Output()

	ctx, cancel := context.WithCancel(context.Background())
	sigChan := make(chan os.Signal, 2)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-sigChan
		cancel()
	}()

	pc, err := newPubsubClient(ctx, pi)
	if err != nil {
		log.Fatalf("failed to create pubsub client: %s", err)
	}
	defer func(pc *pubsubClient) {
		if err := pc.Close(); err != nil {
			log.Printf("closing pubsub client: %s", err)
		}
	}(pc)

	var wg sync.WaitGroup
	brightnessChans := map[uint8]chan<- uint8{}
	reportChans := map[uint8]chan<- struct{}{}
	for id, l := range config.Lights {
		if l.Pi == pi {
			c := make(chan uint8, 8)
			brightnessChans[l.GPIO] = c
			rc := make(chan struct{}, 8)
			reportChans[l.GPIO] = rc
			pin := rpio.Pin(l.GPIO)
			pin.Output()
			wg.Add(1)
			go func(ctx context.Context, id config.ID, pin rpio.Pin, c <-chan uint8, rc <-chan struct{}) {
				pwm(ctx, pc, id, pin, c, rc)
				wg.Done()
			}(ctx, id, pin, c, rc)
		}
	}
	// enable dance mat controls in living room
	if pi == config.RaspiLight {
		go danceMatInput(ctx, brightnessChans)
	}
	wg.Add(1)
	go func(ctx context.Context, chans map[uint8]chan<- uint8) {
		err := pc.ReceiveExecute(ctx, func(ctx context.Context, msg *protocol.ExecuteMessage) {
			for _, cmd := range msg.Commands {
				for _, id := range cmd.Devices {
					light := config.Lights[id]
					if light == nil {
						log.Printf("got message for unknown device: %d", id)
						return
					}
					if light.Pi != pi {
						log.Printf("got message for device %d from different pi: %d (want %d)", id, light.Pi, pi)
						return
					}
					c := chans[light.GPIO]
					if c == nil {
						log.Printf("got message for unknown gpio %d of light %d", light.GPIO, id)
						return
					}
					for _, ex := range cmd.Executions {
						switch ex := ex.(type) {
						case protocol.ExecuteExecutionOnOff:
							if ex.On {
								c <- 255
							} else {
								c <- 0
							}
						default:
							log.Printf("got message with execution of unhandled type %T", ex)
							return
						}
					}
				}
			}
			log.Printf("received message %v", msg)
			// leave reporting back to pwm(), for now
		})
		if err != nil {
			log.Printf("fatal error receiving execute requests: %s", err)
			cancel()
		}
		wg.Done()
	}(ctx, brightnessChans)
	wg.Add(1)
	go func(ctx context.Context, chans map[uint8]chan<- struct{}) {
		err := pc.ReceiveReport(ctx, func(ctx context.Context) {
			for _, c := range chans {
				c <- struct{}{}
			}
			log.Println("received report message")
		})
		if err != nil {
			log.Printf("fatal error receiving report requests: %s", err)
			cancel()
		}
		wg.Done()
	}(ctx, reportChans)
	log.Println("started listening for messages")
	<-ctx.Done()
	wg.Wait()
	for gpio := range brightnessChans {
		rpio.Pin(gpio).Low()
	}
}

func danceMatInput(ctx context.Context, chans map[uint8]chan<- uint8) {
	const (
		jsid     = 0
		pollRate = 10 * time.Millisecond
		errSleep = 5 * time.Second

		btnUp    = 0b0000000000000001 // also Axis0-
		btnDown  = 0b0000000000000010 // also Axis1+
		btnLeft  = 0b0000000000000100 // also Axis1-
		btnRight = 0b0000000000001000 // also Axis0+

		chanRed   = 17
		chanGreen = 22
		chanBlue  = 27
	)
	joy, err := joystick.Open(jsid)
	if err != nil {
		log.Printf("Failed to open joystick: %s", err)
		return
	}
	type State struct {
		up    bool
		down  bool
		left  bool
		right bool
	}
	type Color struct {
		red, green, blue uint8
	}
	prevState := State{}
	prevColor := Color{}
	for {
		rawState, err := joy.Read()
		if err != nil {
			log.Printf("Failed to read joystick: %s", err)
			return
		}
		state := State{
			up:    rawState.Buttons&btnUp != 0,
			down:  rawState.Buttons&btnDown != 0,
			left:  rawState.Buttons&btnLeft != 0,
			right: rawState.Buttons&btnRight != 0,
		}
		color := prevColor
		if state.left && !prevState.left {
			color.red = 255 - color.red
		}
		if state.up && !prevState.up {
			color.green = 255 - color.green
		}
		if state.right && !prevState.right {
			color.blue = 255 - color.blue
		}
		if state.down && !prevState.down {
			color.red = 255 - color.red
			color.green = 255 - color.green
			color.blue = 255 - color.blue
		}
		if color.red != prevColor.red {
			chans[chanRed] <- color.red
		}
		if color.green != prevColor.green {
			chans[chanGreen] <- color.green
		}
		if color.blue != prevColor.blue {
			chans[chanBlue] <- color.blue
		}
		prevColor = color
		prevState = state

		select {
		case <-ctx.Done():
			return
		case <-time.After(pollRate):
		}
	}
}

func pwm(ctx context.Context, pc *pubsubClient, id config.ID, pin rpio.Pin, cmdChan <-chan uint8, reportChan <-chan struct{}) {
	// 50 Hz, each divided into 256 slots -> 12'570 Hz ≈ 20'000 Hz -> 0.05ms
	const tick = 5 * time.Microsecond
	brightness := uint8(0)
	report := func() {
		err := pc.State(ctx, id, protocol.DeviceStates{
			OnOff: &protocol.OnOffState{
				On: brightness > 0, // TODO: dimming
			},
		})
		if err != nil {
			log.Printf("error reporting state back: %s", err)
		}
	}
	for {
		if brightness == 0 {
			pin.Low()
			select {
			case <-ctx.Done():
				return
			case cmd := <-cmdChan:
				brightness = cmd
				report()
				continue
			case <-reportChan:
				report()
				continue
			}
		}
		pin.High()
		select {
		case <-ctx.Done():
			return
		case cmd := <-cmdChan:
			brightness = cmd
			report()
			continue
		case <-reportChan:
			report()
			continue
		case <-time.After(time.Duration(brightness) * tick):
		}
		if brightness < 255 {
			pin.Low()
			select {
			case <-ctx.Done():
				return
			case cmd := <-cmdChan:
				brightness = cmd
				report()
				continue
			case <-reportChan:
				report()
				continue
			case <-time.After(time.Duration(255-brightness) * tick):
				pin.Low()
			}
		}
	}
}
