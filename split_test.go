
package palindrome

import (
	"fmt"
	"testing"
)

func Test_Split(t *testing.T) {
	for i := 1; i <= 9; i++ {
		l, m, r, _, c := Split(i)

		if l != 0 || r != 0 || m != i || c != 0{
			t.Error(fmt.Sprintf("Split error on single digit number, %d %d %d %d %d", i, l, m, r, c))
		}
	}
	for i := 10; i <= 99; i++ {
		lx := i / 10
		rx := i % 10
		l, m, r, _, c := Split(i)

		if l != lx || r != rx || m != -1 || c != 1 {
			t.Error(fmt.Sprintf("Split error on double digit number, %d %d %d %d %d", i, l, m, r, c))
		}
	}
	for i := 100; i <= 999; i++ {
		lx := i / 100
		mx := (i / 10) % 10
		rx := i % 10
		l, m, r, _, c := Split(i)

		if l != lx || r != rx || m != mx || c != 1 {
			t.Error(fmt.Sprintf("Split error on triple digit number, %d %d %d %d %d", i, l, m, r, c))
		}
	}
	for i := 1000; i <= 9999; i += 123 {
		lx := i / 100
		rx := ReverseDigits(i % 100, 2)
		l, m, r, _, c := Split(i)

		if l != lx || r != rx || m != -1 || c != 2 {
			t.Error(fmt.Sprintf("Split error on four digit number, %d %d(%d) %d %d(%d) %d", i, l, lx, m, r, rx, c))
		}
	}
	for i := 10000; i <= 99999; i += 123 {
		lx := i / 1000
		mx := (i / 100) % 10
		rx := ReverseDigits(i % 100, 2)
		l, m, r, _, _ := Split(i)

		if l != lx || r != rx || m != mx {
			t.Error(fmt.Sprintf("Split error on five digit number, %d %d(%d) %d(%d) %d(%d)", i, l, lx, m, mx, r, rx))
		}
	}
	for i := 100000; i <= 999999; i += 1234 {
		lx := i / 1000
		rx := ReverseDigits(i % 1000, 3)
		l, m, r, _, _ := Split(i)

		if l != lx || r != rx || m != -1 {
			t.Error(fmt.Sprintf("Split error on six digit number, %d %d(%d) %d %d(%d)", i, l, lx, m, r, rx))
		}
	}
	for i := 1000000; i <= 9999999; i += 12345 {
		lx := i / 10000
		mx := (i / 1000) % 10
		rx := ReverseDigits(i % 1000, 3)
		l, m, r, _, _ := Split(i)

		if l != lx || r != rx || m != mx {
			t.Error(fmt.Sprintf("Split error on seven digit number, %d %d(%d) %d(%d) %d(%d)", i, l, lx, m, mx, r, rx))
		}
	}
}
