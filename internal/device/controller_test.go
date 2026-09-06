package device_test

import (
	"testing"
	"testing/synctest"

	"github.com/andygeiss/esp32-controller/serial"
	"github.com/andygeiss/esp32-controller/wifi"
	"github.com/andygeiss/esp32/internal/device"
)

// The esp32-controller packages keep the simulated radio and serial port in
// package-level variables, so these tests share state and must not run in
// parallel. They run inside a synctest bubble because Setup waits a second
// between connection attempts and Loop sleeps for five.

func TestSetupConnectsToWifi(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		wifi.CurrentStatus = wifi.StatusIdle
		serial.Baud = 0

		if err := device.Setup(); err != nil {
			t.Fatalf("Setup: %v", err)
		}

		if got, want := wifi.Status(), wifi.StatusConnected; got != want {
			t.Errorf("wifi status: got %v, want %v", got, want)
		}
		if got, want := serial.Baud, serial.BaudRate115200; got != want {
			t.Errorf("baud rate: got %v, want %v", got, want)
		}
	})
}

func TestLoopReturnsNoError(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		if err := device.Setup(); err != nil {
			t.Fatalf("Setup: %v", err)
		}

		if err := device.Loop(); err != nil {
			t.Errorf("Loop: got %v, want nil", err)
		}
	})
}

// The simulator runs NewController, the board runs the transpiled functions.
// This holds the two together: a wrapper that called something else, or
// nothing, would pass every test above.
func TestNewControllerRunsTheTranspiledFunctions(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		wifi.CurrentStatus = wifi.StatusIdle
		serial.Baud = 0
		c := device.NewController()

		if err := c.Setup(); err != nil {
			t.Fatalf("Setup: %v", err)
		}
		if err := c.Loop(); err != nil {
			t.Fatalf("Loop: %v", err)
		}

		// Only device.Setup sets the baud rate, so a wrapper that skipped it
		// leaves this at zero.
		if got, want := serial.Baud, serial.BaudRate115200; got != want {
			t.Errorf("baud rate: got %v, want %v", got, want)
		}
		if got, want := wifi.Status(), wifi.StatusConnected; got != want {
			t.Errorf("wifi status: got %v, want %v", got, want)
		}
	})
}
