package serial_test

import (
	"testing"

	. "github.com/andygeiss/assert"

	"github.com/andygeiss/esp32/business/controller/serial"
)

func TestSerialBegin(t *testing.T) {
	baud := serial.BaudRate115200
	serial.Baud = 0
	serial.Begin(baud)
	Assert(t, serial.Baud, IsEqual(serial.BaudRate115200))
}
