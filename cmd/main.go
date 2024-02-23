package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

// each line, one or more numerical chars
// get he first and the last
// the first may also be the last

//go:embed input1.txt
var inputDayOne string

func main() {
	fmt.Println("hello world")
	// fmt.Println(inputDayOne)
	lines := strings.Split(strings.Trim(inputDayOne, "\n"), "\n")
	fmt.Println(SumLines(lines))
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return result
}

func FirstNumber(line string) string {
	pos := strings.IndexAny(line, "123456789")
	if pos != -1 {
		return string(line[pos])
	}
	return "NOTFOUND"
}

func LastNumber(line string) string {
	rl := reverse(line)
	pos := strings.IndexAny(rl, "123456789")
	if pos != -1 {
		return string(rl[pos])
	}
	return "NOTFOUND"
}

func FirstLast(line string) string {
	return FirstNumber(line) + LastNumber(line)
}

func SumLines(lines []string) int {
	total := 0
	for _, l := range lines {
		n, err := strconv.Atoi(FirstLast(l))
		if err != nil {
			fmt.Printf("Could not convert %v to number\n", l)
		}
		total += n
	}
	return total
}
