package device_test

import (
	"testing"

	"github.com/andygeiss/cloud-native-utils/assert"
	"github.com/andygeiss/esp32/device"
)

func TestControllerSetupErrorShouldBeNil(t *testing.T) {
	ctrl := device.NewController()
	err := ctrl.Setup()
	assert.That(t, "setup should return without an error", err, nil)
}

func TestControllerLoopErrorShouldBeNil(t *testing.T) {
	ctrl := device.NewController()
	ctrl.Setup()
	err := ctrl.Loop()
	assert.That(t, "loop once should return without an error", err, nil)
}
