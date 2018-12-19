package device_test

import (
	"github.com/andygeiss/assert"
	"github.com/andygeiss/assert/is"
	"github.com/andygeiss/esp32/device"
	"testing"
)

func TestControllerSetupErrorShouldBeNil(t *testing.T) {
	ctrl := device.NewController()
	err := ctrl.Setup()
	assert.That(t, err, is.Equal(nil))
}

func TestControllerLoopErrorShouldBeNil(t *testing.T) {
	ctrl := device.NewController()
	ctrl.Setup()
	err := ctrl.Loop()
	assert.That(t, err, is.Equal(nil))
}
