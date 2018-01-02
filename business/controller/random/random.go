package random

import "math/rand"

// Num returns pseudo-random numbers.
// @see: https://www.arduino.cc/reference/en/language/functions/random-numbers/random/
func Num(max int) int {
	return rand.Intn(max)
}

// NumBetween returns pseudo-random numbers between min and max.
// @see: https://www.arduino.cc/reference/en/language/functions/random-numbers/random/
func NumBetween(min, max int) int {
	return rand.Intn(min) + (max - min)
}

// Seed initializes the pseudo-random number generator, causing it to start at an arbitrary point in its random sequence.
// @see: https://www.arduino.cc/reference/en/language/functions/random-numbers/randomseed/
func Seed(seed int64) {
	rand.Seed(seed)
}
