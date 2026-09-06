// Package simulator runs a controller the way an ESP32 runs it: set it up once,
// then loop it.
package simulator

import (
	"context"
	"fmt"

	controller "github.com/andygeiss/esp32-controller"
)

// Run sets c up and then loops it. It stops after loops iterations, or when ctx
// is cancelled if loops is 0. A negative loops is an error.
//
// Cancellation is only seen between iterations: the controller sleeps through
// timer.Delay, which takes no context because the Arduino call it stands in for
// has none.
func Run(ctx context.Context, c controller.Controller, loops int) error {
	if loops < 0 {
		return fmt.Errorf("loops %d: want zero or more", loops)
	}
	if err := c.Setup(); err != nil {
		return fmt.Errorf("setup: %w", err)
	}
	for i := 0; loops == 0 || i < loops; i++ {
		// Loops are numbered from one, so the message reads the way a person
		// counts them.
		if err := ctx.Err(); err != nil {
			return fmt.Errorf("interrupted before loop %d: %w", i+1, err)
		}
		if err := c.Loop(); err != nil {
			return fmt.Errorf("loop %d: %w", i+1, err)
		}
	}
	return nil
}
