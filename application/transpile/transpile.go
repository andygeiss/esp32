package transpile

import (
	"fmt"

	"github.com/andygeiss/esp32/business/transpiler"
	"github.com/andygeiss/esp32/business/worker"
)

// Transpiler uses a given worker to transpile source code.
type Transpiler struct {
	w worker.Worker
}

const (
	// ErrorTranspilerWorkerIsNil ...
	ErrorTranspilerWorkerIsNil = "Worker should not be nil"
)

// NewTranspiler creates a new transpiler and returns its address.
func NewTranspiler(w worker.Worker) transpiler.Transpiler {
	return &Transpiler{w}
}

// Transpile invokes the workers
func (c *Transpiler) Transpile() error {
	if c.w == nil {
		return fmt.Errorf("Error: %v", ErrorTranspilerWorkerIsNil)
	}
	return c.w.Start()
}
