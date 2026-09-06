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

		if err := device.NewController().Setup(); err != nil {
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
		ctrl := device.NewController()
		if err := ctrl.Setup(); err != nil {
			t.Fatalf("Setup: %v", err)
		}

		if err := ctrl.Loop(); err != nil {
			t.Errorf("Loop: got %v, want nil", err)
		}
	})
}
