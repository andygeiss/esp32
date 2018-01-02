package digital

const (
	// High ...
	High = 1
	// Low ...
	Low = 0
	// ModeInput ...
	ModeInput = 0
	// ModeOutput ...
	ModeOutput = 1
	// PinsMax ...
	PinsMax = 48
)

var (
	// GPIOModes ...
	GPIOModes = make(map[int]int, PinsMax)
	// GPIOValues ...
	GPIOValues = make(map[int]int, PinsMax)
)

// IsPinValid ...
func IsPinValid(pin int) bool {
	return pin <= PinsMax && pin > 0
}

// PinMode configures the specified pin to behave either as an input or an output.
// @see: https://www.arduino.cc/reference/en/language/functions/digital-io/pinmode/
func PinMode(pin, mode int) {
	if IsPinValid(pin) {
		GPIOModes[pin] = mode
	}
}

// Write sets a HIGH or a LOW value to a digital pin.
// @see: https://www.arduino.cc/reference/en/language/functions/digital-io/digitalwrite/
func Write(pin, value int) {
	if IsPinValid(pin) {
		GPIOValues[pin] = value
	}
}
