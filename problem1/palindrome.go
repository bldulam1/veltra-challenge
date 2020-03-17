package problem1

func isAlphaNumeric(r rune) bool {
	return ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') || ('0' <= r && r <= '9')
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
	if len(s) == 0 {
		return false
	} else if len(s) == 1 {
		return isAlphaNumeric(rune(s[0]))
	}

	leftPtr, rightPtr := 0, len(s)-1
	isContainAlphaNumeric := false

	for leftPtr < rightPtr {
		for leftPtr < rightPtr && !isAlphaNumeric(rune(s[leftPtr])) {
			leftPtr++
		}
		for leftPtr < rightPtr && !isAlphaNumeric(rune(s[rightPtr])) {
			rightPtr--
		}
		if !isContainAlphaNumeric &&
			(isAlphaNumeric(rune(s[leftPtr])) ||
				isAlphaNumeric(rune(s[rightPtr]))) {
			isContainAlphaNumeric = true
		}

		leftChar, rightChar := toLower(rune(s[leftPtr])), toLower(rune(s[rightPtr]))
		if leftChar != rightChar {
			return false
		}

		leftPtr++
		rightPtr--
	}

	return isContainAlphaNumeric
}

func CheckPalindrome(s string) string {
	if isPalindrome(s) {
		return "A palindrome"
	}
	return "Not a palindrome"
}
