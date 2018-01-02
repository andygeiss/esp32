package timer

import "time"

// Delay pauses the program for the amount of time (in milliseconds) specified as parameter.
// @see: https://www.arduino.cc/reference/en/language/functions/time/delay/
func Delay(ms int) {
	done := make(chan bool)
	go func(ms int) {
		time.Sleep(time.Millisecond * time.Duration(ms))
		done <- true
	}(ms)
	select {
	case <-done:
	}
}
