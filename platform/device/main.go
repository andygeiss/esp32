package main

import (
	"log"

	"github.com/andygeiss/esp32-mqtt/application/device"
)

func main() {
	ctrl := device.NewController()
	if err := ctrl.Setup(); err != nil {
		log.Fatal(err)
	}
	for {
		if err := ctrl.Loop(); err != nil {
			log.Fatal(err)
		}
	}
}
