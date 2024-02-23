package day1_test

import (
	"testing"
"strings"
	"juxt.pro/hackathon/internal/day1"
)

const word_numbers = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func find_lowest_word(str string) (int,string){
    return 0, "one"
}

func to_numeral(word string) (string) {
   return "1"
}

func xxx (line string) (string) {
	lowest_word_index, lowest_word := find_lowest_word(line)
		first_index, first_numeral := day1.FirstNumber(line)
		
		if first_index < lowest_word_index {
		  return first_numeral
		}
		  
		return to_numeral(lowest_word)

}

func TestIndexAny(t *testing.T) {
        //numeral->val := {"one" 1}
		
        line := "xxone9fo4sd9ij2fdsa7one"
		
		numeral := xxx(line)

        //firstpos := strings.Index(line, "one")
        //lastpos := strings.LastIndex(line, "one")
       t.Fatalf("%v ", numeral)
        //expectation := "1"
        //result := day1.FirstNumber(line)
        //if result != expectation {
        //        t.Fatalf("that is wrong, expected %s, got %s", expectation, result)
        //}
}

func TestGetFirstNumber(t *testing.T) {
	line := "9fo4sd9ij2fdsa7"
	expectation := "9"
	_, result := day1.FirstNumber(line)
	if result != expectation {
		t.Fatalf("that is wrong, expected %s, got %s", expectation, result)
	}
}

func TestGetLastNumber(t *testing.T) {
	line := "9fo4sd9ij2fdsa7"
	expectation := "7"
	_, result := day1.LastNumber(line)
	if result != expectation {
		t.Fatalf("that is wrong, expected %s, got %s", expectation, result)
	}
}

func TestFirstLast(t *testing.T) {
	line := "9fo4sd9ij2fdsa7"
	expectation := "97"
	result := day1.FirstLast(line)
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
	result := day1.SumLines(lines)
	if result != expectation {
		t.Fatalf("that is wrong, expected %v, got %v", expectation, result)
	}
}
