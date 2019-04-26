
# go-palindrome #

[![GoDoc](https://godoc.org/github.com/tjmerritt/go-palindrom?status.svg)](https://godoc.org/github.com/tjmerritt/go-palindrome)

While interviewing prospective candidates at three different companies, I have used the following interview question.

Write a function that takes an integer parameter and returns true if it is a palindrome and false otherwise.

Most candidates will produce something like what is contained in str.go, often mistakes are made in the loop termination
condition and the length of the string.  Less frequent but not uncommon are errors in the initial condition setup.
Things I'm interested in are how they go about analyzing the problem, do they reconize that there are tradeoffs between
an all integer solution and the string solution.  Can they produce a correct solution.  If there are errors in the
initial solution, how much prompting does it take for them to correct the errors.


```go
// This is a version of IsPalindrome that converts the integer to a string first and then does
// character by character comparison to check if the string is a palindrome
func IsPalindromeString(num int) bool {
	str := strconv.FormatInt(int64(num), 10)

	// offset to the begining and end of the string
	first := 0
	last := len(str) - 1

	for first < last {
		// if digits do not match, fail
		if str[first] != str[last] {
			return false
		}

		// move to the next pair of digits
		first++
		last--
	}

	// They all matched
	return true
}
```

Benchmarking go test programs indicates that the string solution takes about three times longer than an all integer solution.  The all integer solution is a bit more complex, but not difficult to write.  Few candicates attempt this solution, and fewer come up with a correct solution.  In fairness, unless a candidate is certain that they can knock out an integer solution, I stear them away from doing so.  An optimal integer solution looks like the following.

```go
func IsPalindromeInt(num int) bool {
        left := num
        right := 0
        mult := 1
        nextMult := 10

        for left >= nextMult {
                mult = nextMult
                nextMult *= 10
                digit := left % 10
                left /= 10
                right = right * 10 + digit
        }

        if left >= mult {
                // Adjust for numbers with an odd number of digits.
                left /= 10
        }

        return left == right
}
```

I almost always ask questions about where the time is spent in either case.  The integer solution is dominated by the cost of the divides and multiplies.  Modern processors pretty well flatten out the loop condition.  The 'if' statement if traversing numbers sequentially is well predicted, but not for random numbers.  It is dependent upon the number of digits in the number, so is not evenly split for all values in a range.  These are trivial details that most candidates would not pick up on, but are important for understanding and optimizing performance software.

An iterative loop to tests values in order to find all palindromes in a range does not perform well in general, since the density of palindromes is low and most checks will be false. Candidates will ideally be able to recognize that there is a more efficient solution and be able to explain how to find the next palindrome in a sequence given a know palindrome.  The functions in gen.go provide functions to find the next integer that is a palindrome given a value that may or may not be a palindrome.  The Prev function does the opposite finding the previous palindrom rather than the next.

time.go has a few functions for timing the performance of the various palindrome functions.
