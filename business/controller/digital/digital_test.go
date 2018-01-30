package digital_test

import (
	"testing"

	. "github.com/andygeiss/assert"

	"github.com/andygeiss/esp32/business/controller/digital"
)

func TestDigitalWrite(t *testing.T) {
	pin := 1
	digital.GPIOValues[pin] = digital.Low
	digital.Write(pin, digital.High)
	Assert(t, digital.GPIOValues[pin], IsEqual(digital.High))
}

func TestPinMode(t *testing.T) {
	pin := 1
	digital.GPIOModes[pin] = digital.ModeInput
	digital.PinMode(pin, digital.ModeOutput)
	Assert(t, digital.GPIOModes[pin], IsEqual(digital.ModeOutput))
}
