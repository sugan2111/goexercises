package main

import (
	"testing"
)

func TestFunction1(t *testing.T) {
	c := function1()
	if len(c) != 1 {
		t.Error("Only 1 developer is in 5th floor")
	}
}

func TestFunction3(t *testing.T) {
	c := function3()
	for _,k := range c {
		if (k != "john") {
			t.Error("Expected John, but did not get")
		}
	}
}
