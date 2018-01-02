package device

import (
	"github.com/andygeiss/esp32/business/controller"
	"github.com/andygeiss/esp32/business/controller/serial"
	"github.com/andygeiss/esp32/business/controller/timer"
	wifi "github.com/andygeiss/esp32/business/controller/wifi"
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
	timer.Delay(100)
	serial.Println("Connecting WiFi ...")
	wifi.BeginEncrypted("...SSID...", "...PASSPHRASE...")
	timer.Delay(10000)
	serial.Println(wifi.LocalIP())
	serial.Println("Done.")
	return nil
}
