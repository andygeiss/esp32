package simulator_test

import (
	"context"
	"errors"
	"testing"

	"github.com/andygeiss/esp32/internal/simulator"
)

var errBoom = errors.New("boom")

// fakeController records what the simulator asked it to do, and can cancel the
// run from inside a loop the way a Ctrl-C does.
type fakeController struct {
	setupErr error
	loopErr  error
	stopAt   int
	cancel   context.CancelFunc
	setups   int
	loops    int
}

func (c *fakeController) Setup() error {
	c.setups++
	return c.setupErr
}

func (c *fakeController) Loop() error {
	c.loops++
	if c.cancel != nil && c.loops == c.stopAt {
		c.cancel()
	}
	return c.loopErr
}

func TestRun(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		loops     int
		cancelled bool
		setupErr  error
		loopErr   error
		wantSetup int
		wantLoops int
		wantErr   bool
	}{
		{name: "three loops run three times", loops: 3, wantSetup: 1, wantLoops: 3},
		{name: "one loop runs once", loops: 1, wantSetup: 1, wantLoops: 1},
		{name: "a failing setup stops before the first loop", loops: 3, setupErr: errBoom, wantSetup: 1, wantErr: true},
		{name: "a failing loop stops the run", loops: 3, loopErr: errBoom, wantSetup: 1, wantLoops: 1, wantErr: true},
		{name: "a negative count is rejected", loops: -1, wantErr: true},
		{name: "an interrupted run stops before the first loop", cancelled: true, wantSetup: 1, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx, cancel := context.WithCancel(t.Context())
			defer cancel()
			if tt.cancelled {
				cancel()
			}
			c := &fakeController{setupErr: tt.setupErr, loopErr: tt.loopErr}

			err := simulator.Run(ctx, c, tt.loops)

			if got := err != nil; got != tt.wantErr {
				t.Errorf("error: got %v, want error %v", err, tt.wantErr)
			}
			if got, want := c.setups, tt.wantSetup; got != want {
				t.Errorf("setups: got %d, want %d", got, want)
			}
			if got, want := c.loops, tt.wantLoops; got != want {
				t.Errorf("loops: got %d, want %d", got, want)
			}
		})
	}
}

func TestRunWithoutACountLoopsUntilCancelled(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()
	c := &fakeController{stopAt: 3, cancel: cancel}

	err := simulator.Run(ctx, c, 0)

	if !errors.Is(err, context.Canceled) {
		t.Errorf("error: got %v, want context.Canceled", err)
	}
	if got, want := c.loops, 3; got != want {
		t.Errorf("loops: got %d, want %d", got, want)
	}
}

func TestRunWrapsTheControllerError(t *testing.T) {
	t.Parallel()
	c := &fakeController{loopErr: errBoom}

	err := simulator.Run(t.Context(), c, 1)

	if !errors.Is(err, errBoom) {
		t.Errorf("error: got %v, want it to wrap errBoom", err)
	}
}
