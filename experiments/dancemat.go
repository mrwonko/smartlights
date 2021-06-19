package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/simulatedsimian/joystick"
)

func main() {
	if err := saneMain(os.Args); err != nil {
		log.Fatal(err)
	}
}

func saneMain(args []string) error {
	jsid := 0
	js, err := joystick.Open(jsid)
	if err != nil {
		return fmt.Errorf("opening joystick %d: %w", jsid, err)
	}
	defer js.Close()
	log.Printf("joy %d name: %s", jsid, js.Name())
	numAxes := js.AxisCount()
	log.Printf("axes: %d", numAxes)
	numButtons := js.ButtonCount()
	log.Printf("buttons: %d", numButtons)

	prev := joystick.State{
		AxisData: make([]int, numAxes),
	}
	for {
		cur, err := js.Read()
		if err != nil {
			return fmt.Errorf("reading joystick: %w", err)
		}
		axesDiffer := false
		for i := 0; i < numAxes; i++ {
			if cur.AxisData[i] != prev.AxisData[i] {
				axesDiffer = true
				break
			}
		}
		if cur.Buttons != prev.Buttons || axesDiffer {
			log.Printf(fmt.Sprintf("new state: buttons=%%0%db axes=%%v", numButtons), cur.Buttons, cur.AxisData[:numAxes])
		}
		prev = cur
		time.Sleep(50 * time.Millisecond)
	}
}

const (
	btnUp     = 0b0000000000000001 // also Axis0-
	btnDown   = 0b0000000000000010 // also Axis1+
	btnLeft   = 0b0000000000000100 // also Axis1-
	btnRight  = 0b0000000000001000 // also Axis0+
	btnX      = 0b0000000001000000
	btnO      = 0b0000000010000000
	btnBack   = 0b0000000100000000
	btnSelect = 0b0000001000000000
)
