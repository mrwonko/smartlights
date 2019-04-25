package benchmark

import (
	"testing"

	"github.com/stianeikeland/go-rpio/v4"
)

/*
func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Sleep(time.Millisecond)
	}
}
*/

func BenchmarkOnOff(b *testing.B) {
	const pin = 17

	// FIXME: can't set pin mode to output in my version yet, but it persists so this hack works
	if err := rpio.Open(); err != nil {
		b.Fatal(err)
	}
	rpio.Pin(pin).Output()
	_ = rpio.Close()

	b.Run("library", func(b *testing.B) {
		if err := rpio.Open(); err != nil {
			b.Fatal(err)
		}
		defer rpio.Close()

		b.ResetTimer() // in case Open() is slow

		for i := 0; i < b.N; i++ {
			rpio.Pin(pin).High()
			rpio.Pin(pin).Low()
		}
	})
	b.Run("singlethreaded-one", func(b *testing.B) {
		g, err := newGPIO()
		if err != nil {
			b.Fatal(err)
		}
		defer g.Close()

		b.ResetTimer() // in case newGPIO() is slow

		for i := 0; i < b.N; i++ {
			g.WritePin(pin, false)
			g.WritePin(pin, true)
		}
	})
	b.Run("singlethreaded-multi", func(b *testing.B) {
		g, err := newGPIO()
		if err != nil {
			b.Fatal(err)
		}
		defer g.Close()

		b.ResetTimer() // in case newGPIO() is slow

		for i := 0; i < b.N; i++ {
			g.WritePins(1<<pin, false)
			g.WritePins(1<<pin, true)
		}
	})
}
