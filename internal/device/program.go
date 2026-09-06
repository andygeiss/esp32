package device

import controller "github.com/andygeiss/esp32-controller"

// program is the Go-only half of this package. The board gets setup() and
// loop() from the transpiled functions in controller.go, so nothing on the
// device needs a type; the simulator does, because it runs any
// controller.Controller.
type program struct{}

// NewController returns the controller in controller.go as the interface the
// simulator runs.
func NewController() controller.Controller {
	return program{}
}

// Loop runs the controller's work once.
func (program) Loop() error {
	return Loop()
}

// Setup prepares the board.
func (program) Setup() error {
	return Setup()
}
