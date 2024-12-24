package main

import (
	"testing"
)

func TestPart01(t *testing.T) {
	result := part01("testinput.txt")
	if result != 11 {
		t.Fail()
	}
}

func TestPart02(t *testing.T) {
	result := part02("testinput.txt")
	if result != 31 {
		t.Fail()
	}
}
