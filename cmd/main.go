package main

import (
	_ "embed"
	"flag"
	"fmt"

	"juxt.pro/hackathon/internal/day1"
)

//go:embed ascii.txt
var ascii string

func main() {
	dayPtr := flag.Int("day", 1, "the given day you want to solve")
	partPtr := flag.Int("part", 1, "the given part of the day you want to solve")

	flag.Parse()

	fmt.Println(ascii)
	fmt.Println("Advent of Code: Go Edition")
	fmt.Println("==========================")

	switch *dayPtr {
	case 1:
		if *partPtr == 1 {
			day1.PartOne()
		} else {
			day1.PartTwo()
		}
	case 2:
		fmt.Println("not yet..")
	default:
		fmt.Println("You shouldn't be here..")
	}
}
