package problem1

func isAlphabet(r rune) bool {
	return ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z')
}

func toLower(r rune) rune {
	if 'A' <= r && r <= 'Z' {
		return r + ('a' - 'A')
	}
	return r
}

/*
 Assumptions
 	Palindrome definition: a word, phrase, or sequence that reads the same backwards as forwards, e.g. madam or nurses run.
 	Valid characters of a palindrome are alphabets only.
 	Non-alphabets will be disregarded during checks
*/
func isPalindrome(s string) bool {
	leftPtr, rightPtr := 0, len(s)-1

	for leftPtr < rightPtr {
		// move to the next character from the left if left character is not an alphabet
		for !isAlphabet(rune(s[leftPtr])) {
			leftPtr++
		}
		// move to the next character from the right if right character is not an alphabet
		for !isAlphabet(rune(s[rightPtr])) {
			rightPtr--
		}

		leftChar, rightChar := toLower(rune(s[leftPtr])), toLower(rune(s[rightPtr]))
		if leftChar != rightChar {
			return false
		}

		leftPtr++
		rightPtr--
	}

	return true
}

func CheckPalindrome(s string) string {
	if isPalindrome(s) {
		return "A palindrome"
	}
	return "Not a palindrome"
}
