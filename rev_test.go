

package palindrome

import (
	"fmt"
	"testing"
)

func Test_ReverseDigits(t *testing.T) {
        r := ReverseDigits(-123, 3)
        if r != -321 {
                t.Error(fmt.Sprintf("ReverseDigits error on negative number, %d %d", -123, r))
        }
	for i := 1; i <= 9; i++ {
		r := ReverseDigits(i, 1)

		if r != i {
			t.Error(fmt.Sprintf("ReverseDigits error on single digit number, %d %d", i, r))
		}
	}
	for i := 10; i <= 99; i++ {
		r := ReverseDigits(i, 2)
		rx := (i % 10) * 10 + (i / 10) 

		if r != rx {
			t.Error(fmt.Sprintf("ReverseDigits error on double digit number, %d %d(%d)", i, r, rx))
		}
	}
	for i := 100; i <= 999; i++ {
		r := ReverseDigits(i, 3)
		rx := (i % 10) * 100 + ((i / 10) % 10) * 10 + (i / 100) 

		if r != rx {
			t.Error(fmt.Sprintf("ReverseDigits error on triple digit number, %d %d(%d)", i, r, rx))
		}
	}
	for i := 1000; i <= 9999; i += 123  {
		r := ReverseDigits(i, 4)
		rx := (i % 10) * 1000 + ((i / 10) % 10) * 100 + ((i / 100) % 10) * 10 + (i / 1000) 

		if r != rx {
			t.Error(fmt.Sprintf("ReverseDigits error on four digit number, %d %d(%d)", i, r, rx))
		}
	}
}
