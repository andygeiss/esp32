package device

import (
	controller "github.com/andygeiss/esp32-controller"
	"github.com/andygeiss/esp32-controller/serial"
	wifi "github.com/andygeiss/esp32-controller/wifi"
	"github.com/andygeiss/esp32-controller/timer"
)

// Controller handles the api logic and state of an ESP32.
type Controller struct {
}

const host string = "www.google.com"
var client wifi.Client

// NewController creates a new controller and returns its address.
func NewController() controller.Controller {
	return &Controller{}
}

// Loop code will be called repeatedly.
func (c *Controller) Loop() error {
	serial.Print("Connecting to ")
	serial.Print(host)
	serial.Print(" ...")
	if (client.Connect(host, 443) == wifi.StatusConnected) {
		serial.Println(" Connected!")
		client.Write("GET / HTTP/1.1\r\n")
		client.Write("Host: ")
		client.Write(host)
		client.Write("\r\n\r\n\r\n")
	} else {
		serial.Println(" Failed!")
	}
	return nil
}

// Setup code will be called once.
func (c *Controller) Setup() error {
	serial.Begin(serial.BaudRate115200)
	serial.Print("Connecting to WiFi ...")
	wifi.BeginEncrypted("SSID", "PASS")
	for wifi.Status() != wifi.StatusConnected {
		serial.Print(".")
		timer.Delay(1000)
	}
	serial.Println(" Connected!")
	return nil
}
