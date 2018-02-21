
package palindrome

// Check if a number is a palindrome (base 10)
func Check(n int) bool {
	l, _, r, _, _ := Split(n)
	// If the left side equals the reverse of the right side, we have a palindrom
	// otherwise, we don't.
	return l == r
}

func IsPalindromeInt(n int) bool {
	return Check(n)
}
