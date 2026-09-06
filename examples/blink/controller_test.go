package blink_test

import (
	"testing"
	"testing/synctest"

	"github.com/andygeiss/esp32-controller/digital"
	"github.com/andygeiss/esp32/examples/blink"
)

// The esp32-controller packages keep the simulated pins in package-level maps,
// so these tests share state and must not run in parallel. They run inside a
// synctest bubble because a loop sleeps for a second of wall time.

func TestSetupConfiguresPin2AsOutput(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		digital.GPIOModes[2] = digital.ModeInput

		if err := blink.NewController().Setup(); err != nil {
			t.Fatalf("Setup: %v", err)
		}

		if got, want := digital.GPIOModes[2], digital.ModeOutput; got != want {
			t.Errorf("pin 2 mode: got %v, want %v", got, want)
		}
	})
}

func TestLoopLeavesPin2Low(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctrl := blink.NewController()
		if err := ctrl.Setup(); err != nil {
			t.Fatalf("Setup: %v", err)
		}
		digital.Write(2, digital.High)

		if err := ctrl.Loop(); err != nil {
			t.Fatalf("Loop: %v", err)
		}

		if got, want := digital.GPIOValues[2], digital.Low; got != want {
			t.Errorf("pin 2 value: got %v, want %v", got, want)
		}
	})
}
