package main

import (
	"github.com/andygeiss/esp32/application/device"
	"github.com/andygeiss/esp32/business/controller"
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
