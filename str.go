
package palindrome

import (
	"strconv"
)

func IsPalindromeString(n int) bool {
	s := strconv.FormatInt(int64(n), 10)

	// offset to the begining and end of the string
	b := 0
	e := len(s) - 1

	for b < e {
		// if digits do not match, fail
		if s[b] != s[e] {
			return false
		}

		// move to the next pair of digits
		b++
		e--
	}

	// They all matched
	return true
}
