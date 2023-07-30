package main

import "testing"

func TestSum(t *testing.T) {
	result := sum(2, 3)

	if result != 5 {
		t.Error("The result must be 5")
	}
}
func TestSub(t *testing.T) {
	result := sub(3, 2)

	if result != 1 {
		t.Error("The result must be 1")
	}
}
func TestMult(t *testing.T) {
	result := mult(2, 3)

	if result != 6 {
		t.Error("The result must be 6")
	}
}
func TestDiv(t *testing.T) {
	result := div(6, 3)

	if result != 2 {
		t.Error("The result must be 2")
	}
}

func TestPow(t *testing.T) {
	result := pow(2, 3)

	if result != 8 {
		t.Error("The result must be 8")
	}
}
