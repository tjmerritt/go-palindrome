
package palindrome

// Split a number into a left and right hand portion and a 
// possible middle digit if the number has an odd number
// of digits.  The digits in the right half are returned in
// reverse order.  Thus the number is a palindrome if the 
// left and right halves are equal.
func Split(n int) (left, middle, right, mult, digs int) {
	left = n
	right = 0
	mult = 1
	nextMult := 10
	digs = 0

	for left >= nextMult {
		mult = nextMult
		nextMult *= 10
		digs++
		d := left % 10
		left /= 10
		right = right * 10 + d
	}

	if left >= mult {
		middle = left % 10
		left /= 10
	} else {
		middle = -1
	}

	return
}
