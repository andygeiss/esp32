// Package blink flashes the LED on pin 2, twice a second.
//
// It is a starter controller: copy this file over
// internal/device/controller.go and change its package line to device. It holds
// nothing but what the transpiler reads, so the copy needs no other edit.
package blink

import (
	"github.com/andygeiss/esp32-controller/digital"
	"github.com/andygeiss/esp32-controller/serial"
	"github.com/andygeiss/esp32-controller/timer"
)

// Loop code will be called repeatedly.
func Loop() error {
	timer.Delay(500)
	serial.Println("  Write PIN 2 -> HIGH")
	digital.Write(2, digital.High)
	timer.Delay(500)
	digital.Write(2, digital.Low)
	serial.Println("  Write PIN 2 -> LOW")
	return nil
}

// Setup code will be called once.
func Setup() error {
	serial.Begin(serial.BaudRate115200)
	serial.Println("Setting up PIN 2 -> OUTPUT")
	digital.PinMode(2, digital.ModeOutput)
	serial.Println("Done.")
	return nil
}
