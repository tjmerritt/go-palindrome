
package palindrome

import (
	"testing"
)


type intTestCase struct {
	n int
	palindrome bool
}

var int_tests []intTestCase = []intTestCase{
	intTestCase{ 1, true },
	intTestCase{ 10, false },
	intTestCase{ 11, true },
	intTestCase{ 101, true },
	intTestCase{ 102, false },
	intTestCase{ 1000, false },
	intTestCase{ 1001, true },
	intTestCase{ 1002, false },
}

func Test_IsPalindromeInt(t *testing.T) {
	for _, tc := range int_tests {
		ok := IsPalindromeInt(tc.n)

		if ok != tc.palindrome {
			if tc.palindrome {
				t.Error("IsPalindromeInt %d is a palindrome, but was not identified as one", tc.n)
			} else {
				t.Error("IsPalindromeInt %d is not a palindrome, but was identified as one", tc.n)
			}
		}
	}
}
