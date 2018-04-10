package device

import (
	controller "github.com/andygeiss/esp32-controller"
	"github.com/andygeiss/esp32-controller/digital"
	"github.com/andygeiss/esp32-controller/serial"
	"github.com/andygeiss/esp32-controller/timer"
)

// Controller handles the api logic and state of an ESP32.
type Controller struct {
}

// NewController creates a new controller and returns its address.
func NewController() controller.Controller {
	return &Controller{}
}

// Loop code will be called repeatedly.
func (c *Controller) Loop() error {
	timer.Delay(500)
	serial.Println("  Write PIN 2 -> HIGH")
	digital.Write(2, digital.High)
	timer.Delay(500)
	digital.Write(2, digital.Low)
	serial.Println("  Write PIN 2 -> LOW")
	return nil
}

// Setup code will be called once.
func (c *Controller) Setup() error {
	serial.Begin(serial.BaudRate115200)
	serial.Println("Setting up PIN 2 -> OUTPUT")
	digital.PinMode(2, digital.ModeOutput)
	serial.Println("Done.")
	return nil
}
