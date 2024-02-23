package day1

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input1-henry.txt
var inputDayOne string

func PartOne() {
	fmt.Println("part one!")
	lines := strings.Split(strings.Trim(inputDayOne, "\n"), "\n")
	fmt.Println(SumLines(lines))
}

func PartTwo() {
	fmt.Println("part two!")
	//lines := strings.Split(strings.Trim(inputDayOne, "\n"), "\n")
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return result
}

func FirstNumber(line string) (int,string) {
	pos := strings.IndexAny(line, "123456789")
	if pos != -1 {
		return pos, string(line[pos])
	}
	return pos, "NOTFOUND"
}

func LastNumber(line string) (int, string) {
	rl := reverse(line)
	pos := strings.IndexAny(rl, "123456789")
	if pos != -1 {
		return pos, string(rl[pos])
	}
	return pos, "NOTFOUND"
}

func FirstLast(line string) string {
	_, f := FirstNumber(line)
    _, l := LastNumber(line)
	return  f + l 
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
