package random_test

import (
	"testing"

	"github.com/andygeiss/esp32/business/controller/random"
)

func TestNum(t *testing.T) {
	if num := random.Num(10); num <= 0 {
		t.Errorf("Num should be greater than 0! %v", num)
	}
}

func TestNumXY(t *testing.T) {
	x := random.Num(100)
	y := random.Num(100)
	if x == y {
		t.Error("X and Y should not be the same!")
	}
}

func TestNumBetween(t *testing.T) {
	num := random.NumBetween(100, 200)
	if num < 100 || num > 200 {
		t.Errorf("Num should be between 100 and 200! %v", num)
	}
}

func TestSeed(t *testing.T) {
	random.Seed(42)
	x := random.Num(100)
	y := random.Num(100)
	if x == y {
		t.Error("X and Y should not be the same!")
	}
}
