
package palindrome

// Return the number that has the reserve of the last
// digits of n.  If n is negative the return must be negative
func ReverseDigits(n, digs int) int {
	if n < 0 {
		return -ReverseDigits(-n, digs)
	}

	r := 0

	for d := 1; d <= digs; d++ {
		// remove one digit from n and add it to r
		r = r * 10 + n % 10
		n /= 10
	}

	return r
}
