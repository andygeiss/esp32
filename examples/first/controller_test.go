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
	err := ctrl.Setup()
	assert.That(t, "setup should return without an error", err, nil)
	err = ctrl.Loop()
	assert.That(t, "loop once should return withoout an error", err, nil)
}
