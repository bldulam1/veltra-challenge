package main

import (
	"fmt"
	"veltra-challenge/problem1"
	"veltra-challenge/problem2"
)

func p2TestDriver() {
	years := []int{
		-1, 0, 4, 10, 20, 100, 1700, 1800, 1900, 1600, 2000, 2019, 2020,
	}

	for _, year := range years {
		fmt.Println(year, ":", problem2.IsLeapYear(year))
	}
}

func p1TestDriver() {
	words := []string{
		"",
		"?",
		"0",
		"A",
		"AB",
		"11",
		"11a",
		"113",
		"Hello World",
		"A Toyota’s a Toyota",
		"Nurses run",
	}

	for _, word := range words {
		fmt.Println(problem1.CheckPalindrome(word), ":", word)
	}
}

func main() {
	fmt.Println("Problem 1")
	p1TestDriver()

	fmt.Println()

	fmt.Println("Problem 2")
	p2TestDriver()
}
