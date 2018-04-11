package main

import (
	"fmt"
	"github.com/andygeiss/esp32/device"
	"os"
)

func main() {
	ctrl := device.NewController()
	if err := ctrl.Setup(); err != nil {
		fmt.Fprintf(os.Stderr, "Error on Setup: %s", err.Error())
	}
	for {
		if err := ctrl.Loop(); err != nil {
			fmt.Fprintf(os.Stderr, "Error on Loop: %s", err.Error())
		}
	}
}
