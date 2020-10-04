package device_test

import (
	"testing"

	"github.com/andygeiss/esp32/device"
	"github.com/andygeiss/esp32/pkg/assert"
)

func TestControllerSetupErrorShouldBeNil(t *testing.T) {
	ctrl := device.NewController()
	err := ctrl.Setup()
	assert.That("setup should return without an error", t, err, nil)
}

func TestControllerLoopErrorShouldBeNil(t *testing.T) {
	ctrl := device.NewController()
	ctrl.Setup()
	err := ctrl.Loop()
	assert.That("loop once should return without an error", t, err, nil)
}
