// Package device holds the controller that runs on the ESP32. controller.go is
// the file you edit.
//
// Only controller.go is transpiled to an Arduino sketch, and the transpiler
// understands a small subset of Go: plain functions, constants, package-level
// variables and calls into esp32-controller. It refuses anything else by name
// and line rather than pass over it, so a type, a method, a range loop or a
// call into the standard library stops `make firmware` instead of reaching the
// board. Setup and Loop are plain functions here for that reason; program.go
// wraps them in the interface the simulator runs, and the transpiler never
// reads that file.
package device

import (
	"github.com/andygeiss/esp32-controller/serial"
	"github.com/andygeiss/esp32-controller/timer"
	wifi "github.com/andygeiss/esp32-controller/wifi"
)

const ssid string = "SSID"
const pass string = "PASSPHRASE"
const host string = "HOSTNAME"
const port int = 3000
const request string = "GET /index.html HTTP/1.0\r\n\r\n"

var client wifi.Client

// Loop code will be called repeatedly.
func Loop() error {
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
func Setup() error {
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
