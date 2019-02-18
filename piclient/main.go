package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/mrwonko/smartlights/config"

	rpio "github.com/stianeikeland/go-rpio/v4"
)

const (
	red   = rpio.Pin(13)
	green = rpio.Pin(19)
	blue  = rpio.Pin(18)
)

func main() {
	const (
		pi = 0
	)
	err := rpio.Open()
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
	chans := map[uint8]chan<- uint8{}
	for _, l := range config.Lights {
		if l.Pi == pi {
			c := make(chan uint8, 8)
			chans[l.GPIO] = c
			pin := rpio.Pin(l.GPIO)
			pin.Output()
			wg.Add(1)
			go func(ctx context.Context, pin rpio.Pin, c <-chan uint8) {
				pwm(ctx, pin, c)
				wg.Done()
			}(ctx, pin, c)
		}
	}
	wg.Add(1)
	go func(ctx context.Context, chans map[uint8]chan<- uint8) {
		pc.ReceiveExecute(ctx, func(ctx context.Context, msg *config.ExecuteMessage) {
			c := chans[msg.GPIO]
			if c == nil {
				log.Printf("got message for unknown gpio: %v", msg)
				return
			}
			if msg.On {
				c <- 255
			} else {
				c <- 0
			}
			log.Printf("received message %v", msg)
		})
		wg.Done()
	}(ctx, chans)
	log.Println("started listening for messages")
	<-ctx.Done()
	wg.Wait()
	for gpio := range chans {
		rpio.Pin(gpio).Low()
	}
}

func pwm(ctx context.Context, pin rpio.Pin, cmdChan <-chan uint8) {
	// 50 Hz, each divided into 256 slots -> 12'570 Hz ≈ 20'000 Hz -> 0.05ms
	const tick = 5 * time.Microsecond
	brightness := uint8(0)
	for {
		if brightness == 0 {
			pin.Low()
			select {
			case <-ctx.Done():
				return
			case cmd := <-cmdChan:
				brightness = cmd
				continue
			}
		}
		pin.High()
		select {
		case <-ctx.Done():
			return
		case cmd := <-cmdChan:
			brightness = cmd
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
				continue
			case <-time.After(time.Duration(255-brightness) * tick):
				pin.Low()
			}
		}
	}
}
