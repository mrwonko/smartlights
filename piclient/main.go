package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	rpio "github.com/stianeikeland/go-rpio/v4"
)

const (
	red   = rpio.Pin(13)
	green = rpio.Pin(19)
	blue  = rpio.Pin(18)
)

func main() {
	err := rpio.Open()
	if err != nil {
		log.Fatalf("failed to init go-rpio: %s", err)
	}
	defer rpio.Close()

	red.Output()
	green.Output()
	blue.Output()
	redBrightness := make(chan uint8, 8)
	greenBrightness := make(chan uint8, 8)
	blueBrightness := make(chan uint8, 8)

	ctx, cancel := context.WithCancel(context.Background())
	sigChan := make(chan os.Signal, 2)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-sigChan
		cancel()
	}()

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		pwm(ctx, red, redBrightness)
		wg.Done()
	}()
	go func() {
		pwm(ctx, green, greenBrightness)
		wg.Done()
	}()
	go func() {
		pwm(ctx, blue, blueBrightness)
		wg.Done()
	}()
	redBrightness <- 255
	greenBrightness <- 0
	blueBrightness <- 255
	<-ctx.Done()
	wg.Wait()
	red.Low()
	green.Low()
	blue.Low()
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
