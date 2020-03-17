package main

import (
	"fmt"
	"veltra-challenge/problem2"
)

func main() {
	years := []int{
		-1, 0, 4, 10, 20, 100, 1700, 1800, 1900, 1600, 2000, 2019, 2020,
	}

	for _, year := range years {
		fmt.Println(year, ":", problem2.IsLeapYear(year))
	}
}
