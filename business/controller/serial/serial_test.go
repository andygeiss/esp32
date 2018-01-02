package serial_test

import (
	"testing"

	"github.com/andygeiss/goat/business/controller/serial"
)

func TestSerialBegin(t *testing.T) {
	baud := serial.BaudRate115200
	serial.Baud = 0
	serial.Begin(baud)
	if serial.Baud != baud {
		t.Error("Baud rate should be serial.Baudrate115200!")
	}
}
