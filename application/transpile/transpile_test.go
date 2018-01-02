package transpile_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/andygeiss/esp32-mqtt/application/transpile"
	"github.com/andygeiss/esp32-mqtt/business/worker"
)

type mockupWorker struct {
	in  io.Reader
	out io.Writer
}

func (w *mockupWorker) Prepare(source []worker.Source) (chan worker.Source, error) {
	out := make(chan worker.Source)
	return out, nil
}
func (w *mockupWorker) Start() error {
	return nil
}
func (w *mockupWorker) Transform(source chan worker.Source) (chan worker.Target, error) {
	out := make(chan worker.Target)
	return out, nil
}

func TestTranspileShouldBeSuccessful(t *testing.T) {
	var in, out bytes.Buffer
	worker := &mockupWorker{&in, &out}
	trans := transpile.NewTranspiler(worker)
	if err := trans.Transpile(); err != nil {
		t.Errorf("Transpile() should not return an error! [%s]", err.Error())
	}
}
