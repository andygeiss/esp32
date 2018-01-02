package device

import (
	"github.com/andygeiss/esp32/business/controller"
	"github.com/andygeiss/esp32/business/controller/serial"
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
	return nil
}

// Setup code will be called once.
func (c *Controller) Setup() error {
	serial.Begin(serial.BaudRate115200)
	return nil
}
