
package palindrome

func multiplier(n int) int {
	// Calculate 10^n efficiently
	f := 10
	m := 1

	if n < 0 {
		return -1
	}

	for n > 0 {
		if n & 1 != 0 {
			m *= f
		}
		f *= f
		n >>= 1
	}

	return m
}

func Construct(l, m, x, c int) int {
	if x == 0 {
		x = multiplier(c)
	}

	if m >= 0 {
		// Construct a palindrome with an odd number of digits
		return (l * 10 + m) * x + ReverseDigits(l, c)
	}

	// Construct a palindrome with an even number of digits
	return l * x + ReverseDigits(l, c)
}

func Prev(n int) int {
	// Split the number up into the left and right halves
	l, m, r, x, c := Split(n)

	// If the right half is less than or equal to the reverse of left
	// then the previous palindrome has a left side that is one less
	if ReverseDigits(r, c) <= ReverseDigits(l, c) {
		if m >= 0 {
			// We have a middle digit, so decrement it
			m--
			if m < 0 {
				// The middle digits underflowed so decrement
				// the left side and reset the middle digit to 9
				l--
				m = 9
			}
		} else {
			// No middle digit, just decrement the left side
			l--
		}

		if l < x / 10 {
			// The left side underflowed
			if m >= 0 {
				// There was a middle digit, add it to the left
				// side and drop the middle digit
				l = l * 10 + m
				m = -1
			} else {
				// There was no middle digit, so add one and
				// drop a digit off the right side
				m = 9
				c--
				x /= 10
			}
		}
	}

	return Construct(l, m, x, c)
}

func Next(n int) int {
	// Split the number up into the left and right halves
	l, m, r, x, c := Split(n)

	// If the right half is greater than or equal to the reverse of left
	// then the next palindrome has a left side that is one greater
	if ReverseDigits(r, c) >= ReverseDigits(l, c) {
		if m >= 0 {
			// We have a middle digit, so increment it
			m++
			if m > 9 {
				// The middle digits overflowed so increment
				// the left side and reset the middle digit to 0
				l++
				m = 0
			}
		} else {
			// No middle digit, just increment the left side
			l++
		}

		if l >= x {
			// The left side overflowed
			if m >= 0 {
				// There was a middle digit, so just drop
				// it and increment the digit count
				m = -1
				c++
				x *= 10
			} else {
				// There was no middle digit, so add one and
				// drop a digit off the left side
				m = 0
				l /= 10
			}
		}
	}

	return Construct(l, m, x, c)
}
