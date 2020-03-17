package main

import (
	"fmt"
	"veltra-challenge/problem1"
)

func main() {
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
