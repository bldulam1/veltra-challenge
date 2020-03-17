package problem2

/*
	leap year (definition)
	a year, occurring once every four years, which has 366 days including 29 February as an intercalary day.

	reference: https://en.wikipedia.org/wiki/Leap_year
	Every year that is exactly divisible by four is a leap year,
	except for years that are exactly divisible by 100,
	but these centurial years are leap years if they are exactly divisible by 400.

	For example, the years 1700, 1800, and 1900 are not leap years, but the years 1600 and 2000 are.
*/

func isDivisible(num int, divisor int) bool {
	if num%divisor == 0 {
		return true
	}
	return false
}

func IsLeapYear(year int) bool {
	if year < 0 {
		return false
	} else if isDivisible(year, 400) {
		return true
	} else if isDivisible(year, 100) {
		return false
	} else if isDivisible(year, 4) {
		return true
	}
	return false
}
