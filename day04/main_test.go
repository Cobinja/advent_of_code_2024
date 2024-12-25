package main

import (
	"testing"
)

func TestPart01(t *testing.T) {
	result := part01("testinput.txt")
	if result != 18 {
		t.Fail()
	}
}

func TestPart02(t *testing.T) {
	result := part02("testinput.txt")
	if result != 9 {
		t.Fail()
	}
}
