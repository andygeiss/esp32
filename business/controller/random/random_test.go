package random_test

import (
	"testing"

	. "github.com/andygeiss/assert"

	"github.com/andygeiss/esp32/business/controller/random"
)

func TestNum(t *testing.T) {
	Assert(t, random.Num(10), IsNotEqual(0))
}

func TestNumXY(t *testing.T) {
	x := random.Num(100)
	y := random.Num(100)
	Assert(t, x, IsNotEqual(y))
}

func TestNumBetween(t *testing.T) {
	num := random.NumBetween(100, 200)
	Assert(t, num, IsBetween(100, 200))
}

func TestSeed(t *testing.T) {
	random.Seed(42)
	x := random.Num(100)
	y := random.Num(100)
	Assert(t, x, IsNotEqual(y))
}
