package serial

import "fmt"

const (
	// BaudRate300 ...
	BaudRate300 = 300
	// BaudRate600 ...
	BaudRate600 = 600
	// BaudRate1200 ...
	BaudRate1200 = 1200
	// BaudRate2400 ...
	BaudRate2400 = 2400
	// BaudRate4800 ...
	BaudRate4800 = 4800
	// BaudRate9600 ...
	BaudRate9600 = 9600
	// BaudRate14400 ...
	BaudRate14400 = 14400
	// BaudRate19200 ...
	BaudRate19200 = 19200
	// BaudRate28800 ...
	BaudRate28800 = 28800
	// BaudRate38400 ...
	BaudRate38400 = 38400
	// BaudRate57600 ...
	BaudRate57600 = 57600
	// BaudRate115200 ...
	BaudRate115200 = 115200
)

var (
	// AvailableN ...
	AvailableN = 0
	// Baud ...
	Baud = 0
)

// Available gets the number of bytes (characters) available for reading from the serial port.
// @see: https://www.arduino.cc/reference/en/language/functions/communication/serial/available/
func Available() int {
	return AvailableN
}

// Begin sets the data rate in bits per second (baud) for serial data transmission.
// @see: https://www.arduino.cc/reference/en/language/functions/communication/serial/begin/
func Begin(baud int) {
	AvailableN = 1
	Baud = baud
}

// Print prints data to the serial port as human-readable ASCII text.
// @see: https://www.arduino.cc/reference/en/language/functions/communication/serial/print/
func Print(val interface{}) {
	fmt.Print(val)
}

// Println prints data to the serial port as human-readable ASCII text followed by a carriage return character (ASCII 13, or '\r') and a newline character (ASCII 10, or '\n').
// @see: https://www.arduino.cc/reference/en/language/functions/communication/serial/println/
func Println(val interface{}) {
	fmt.Println(val)
}
