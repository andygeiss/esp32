package device

import (
	"github.com/andygeiss/esp32/business/controller"
	"github.com/andygeiss/esp32/business/controller/serial"
	"github.com/andygeiss/esp32/business/controller/timer"
	wifi "github.com/andygeiss/esp32/business/controller/wifi"
	"github.com/andygeiss/esp32/business/controller/digital"
)

// Controller handles the business logic and state of an ESP32.
type Controller struct {
}

// NewController creates a new controller and returns its address.
func NewController() controller.Controller {
	return &Controller{}
}

// Loop code will be called repeatedly.
func (c *Controller) Loop() error {
	timer.Delay(500)
	digital.Write(2, digital.High)
	timer.Delay(500)
	digital.Write(2, digital.Low)
	return nil
}

// Setup code will be called once.
func (c *Controller) Setup() error {
	digital.PinMode(2, digital.ModeOutput)
	return nil
}

