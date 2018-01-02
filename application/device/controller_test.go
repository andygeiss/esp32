package device_test

import (
	"testing"

	"github.com/andygeiss/esp32/application/device"
)

func TestControllerSetupIsSuccessful(t *testing.T) {
	ctrl := device.NewController()
	if err := ctrl.Setup(); err != nil {
		t.Errorf("Setup should not return an error! %v", err)
	}
}

func TestControllerLoopIsSuccessful(t *testing.T) {
	ctrl := device.NewController()
	ctrl.Setup()
	if err := ctrl.Loop(); err != nil {
		t.Errorf("Loop should not return an error! %v", err)
	}
}
