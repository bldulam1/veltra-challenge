package main

import (
	"fmt"
	"veltra-challenge/problem1"
)

func main() {
	words := []string{
		"Hello World",
		"A Toyota’s a Toyota",
	}

	for _, word := range words {
		fmt.Println(problem1.CheckPalindrome(word), ":", word)
	}
}
