package main

import (
	controller "github.com/andygeiss/esp32-controller"
	"github.com/andygeiss/esp32/device"
	"github.com/andygeiss/log"
)

func main() {
	ctrl := device.NewController()
	safeSetup(ctrl)
	for {
		safeLoop(ctrl)
	}
}

func safeLoop(ctrl controller.Controller) {
	if err := ctrl.Loop(); err != nil {
		log.Fatal("%v", err)
	}
}

func safeSetup(ctrl controller.Controller) {
	if err := ctrl.Setup(); err != nil {
		log.Fatal("%v", err)
	}
}
