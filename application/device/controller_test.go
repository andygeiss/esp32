package device_test

import (
	"testing"

	. "github.com/andygeiss/assert"

	"github.com/andygeiss/esp32/application/device"
)

func TestControllerSetupErrorShouldBeNil(t *testing.T) {
	ctrl := device.NewController()
	Assert(t, ctrl.Setup(), IsNil())
}

func TestControllerLoopErrorShouldBeNil(t *testing.T) {
	ctrl := device.NewController()
	ctrl.Setup()
	Assert(t, ctrl.Loop(), IsNil())
}
