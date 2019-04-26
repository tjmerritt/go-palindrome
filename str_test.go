
package palindrome

import (
	"testing"
)


type TestCase struct {
	n int
	palindrome bool
}

var str_tests []TestCase = []TestCase{
	TestCase{ 1, true },
	TestCase{ 10, false },
	TestCase{ 11, true },
	TestCase{ 101, true },
	TestCase{ 102, false },
	TestCase{ 1000, false },
	TestCase{ 1001, true },
	TestCase{ 1002, false },
}

func Test_IsPalindromeString(t *testing.T) {
	for _, tc := range str_tests {
		ok := IsPalindromeString(tc.n)

		if ok != tc.palindrome {
			if tc.palindrome {
				t.Errorf("IsPalindromeString %d is a palindrome, but was not identified as one", tc.n)
			} else {
				t.Errorf("IsPalindromeString %d is not a palindrome, but was identified as one", tc.n)
			}
		}
	}
}
