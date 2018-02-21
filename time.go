
package palindrome

import (
	"time"
)

func TimeLinearInt(min, max int) (count int, t time.Duration) {
	st := time.Now()

	for i := min; i <= max; i++ {
		if IsPalindromeInt(i) {
			count++
		}
	}

	t = time.Now().Sub(st)
	return
}

func TimeLinearStr(min, max int) (count int, t time.Duration) {
	st := time.Now()

	for i := min; i <= max; i++ {
		if IsPalindromeString(i) {
			count++
		}
	}

	t = time.Now().Sub(st)
	return
}

func TimeNextGenerator(min, max int) (count int, t time.Duration) {
	p := min
	
	if p > 9 {
		p = Next(p - 1)
	}

	st := time.Now()

	for p <= max {
		count++
		p = Next(p)
	}

	t = time.Now().Sub(st)
	return
}

func TimePrevGenerator(min, max int) (count int, t time.Duration) {
	p := Prev(max + 1)

	st := time.Now()

	for p >= min {
		count++
		p = Prev(p)
	}

	t = time.Now().Sub(st)
	return
}
