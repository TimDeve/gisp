package main

import "testing"

func TestOne(t *testing.T) {
	if 1 != 1 {
		t.Errorf("Maths are broken")
	}
}
