package device

import (
	controller "github.com/andygeiss/esp32-controller"
	"github.com/andygeiss/esp32-controller/serial"
	"github.com/andygeiss/esp32-controller/timer"
	wifi "github.com/andygeiss/esp32-controller/wifi"
)

// Controller handles the api logic and state of an ESP32.
type Controller struct {
}

const ssid string = "SSID"
const pass string = "PASSPHRASE"
const host string = "HOSTNAME"
const port int = 3000
const request string = "GET /index.html HTTP/1.0\r\n\r\n"

var client wifi.Client

// NewController creates a new controller and returns its address.
func NewController() controller.Controller {
	return &Controller{}
}

// Loop code will be called repeatedly.
func (c *Controller) Loop() error {
	serial.Print("Connecting to [")
	serial.Print(host)
	serial.Print(":")
	serial.Print(port)
	serial.Print(")] ...")
	if client.Connect(host, port) {
		serial.Println(" Connected!")
		client.Println(request)
	} else {
		serial.Println(" Failed!")
	}
	timer.Delay(5000)
	return nil
}

// Setup code will be called once.
func (c *Controller) Setup() error {
	serial.Begin(serial.BaudRate115200)
	serial.Print("Connecting to WiFi ...")
	for wifi.Status() != wifi.StatusConnected {
		wifi.BeginEncrypted(ssid, pass)
		serial.Print(".")
		timer.Delay(1000)
	}
	serial.Println(" Connected!")
	return nil
}
