package main

import (
	"testing"
)

func TestPart01(t *testing.T) {
	result := part01("testinput01.txt")
	if result != 161 {
		t.Fail()
	}
}

func TestPart02(t *testing.T) {
	result := part02("testinput02.txt")
	if result != 48 {
		t.Fail()
	}
}
