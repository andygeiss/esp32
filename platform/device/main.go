package main

import (
	"github.com/andygeiss/log"

	"github.com/andygeiss/esp32/application/device"
)

func main() {
	ctrl := device.NewController()
	if err := ctrl.Setup(); err != nil {
		log.Fatal("%v", err)
	}
	for {
		if err := ctrl.Loop(); err != nil {
			log.Fatal("%v", err)
		}
	}
}
