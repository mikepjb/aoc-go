package main

import (
	"testing"
)

func TestGetFirstNumber(t *testing.T) {
	line := "9fo4sd9ij2fdsa7"
	expectation := "9"
	result := FirstNumber(line)
	if result != expectation {
		t.Fatalf("that is wrong, expected %s, got %s", expectation, result)
	}
}

func TestGetLastNumber(t *testing.T) {
	line := "9fo4sd9ij2fdsa7"
	expectation := "7"
	result := LastNumber(line)
	if result != expectation {
		t.Fatalf("that is wrong, expected %s, got %s", expectation, result)
	}
}

func TestFirstLast(t *testing.T) {
	line := "9fo4sd9ij2fdsa7"
	expectation := "97"
	result := FirstLast(line)
	if result != expectation {
		t.Fatalf("that is wrong, expected %s, got %s", expectation, result)
	}
}

func TestSumLines(t *testing.T) {
	lines := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}
	expectation := 142
	result := SumLines(lines)
	if result != expectation {
		t.Fatalf("that is wrong, expected %v, got %v", expectation, result)
	}
}
