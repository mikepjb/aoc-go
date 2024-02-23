package main

import "testing"

func TestGetFirstNumber(t *testing.T) {
	line := "4fosdijfdsa7"
	if FirstNumber(line) != "4" {
		t.Fatalf("that is wrong")
	}
}
