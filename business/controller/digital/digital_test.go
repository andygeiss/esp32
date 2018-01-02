package digital_test

import (
	"testing"

	"github.com/andygeiss/esp32/business/controller/digital"
)

func TestDigitalWrite(t *testing.T) {
	pin := 1
	digital.GPIOValues[pin] = digital.Low
	digital.Write(pin, digital.High)
	if digital.GPIOValues[pin] != digital.High {
		t.Error("GPIO value should be digitial.High!")
	}
}

func TestPinMode(t *testing.T) {
	pin := 1
	digital.GPIOModes[pin] = digital.ModeInput
	digital.PinMode(pin, digital.ModeOutput)
	if digital.GPIOModes[pin] != digital.ModeOutput {
		t.Error("GPIO mode should be digitial.ModeOutput!")
	}
}
