package main

import (
	_ "embed"
	"fmt"
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
	line := strings.Split(inputDayOne, "\n")
	fmt.Println(line[0])
}

func FirstNumber(line string) string {
	return "4"
}
