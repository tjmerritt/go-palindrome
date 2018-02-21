
package main

import (
	"fmt"

	"github.com/tjmerritt/go-palindrome"
)

func main() {
	min := 1
	max := 10000000
	n1, t1 := palindrome.TimeLinearInt(min, max)
	fmt.Printf("Linear Int %d to %d: count %d: time %v\n", min, max, n1, t1)
	n2, t2 := palindrome.TimeLinearStr(min, max)
	fmt.Printf("Linear Str %d to %d: count %d: time %v\n", min, max, n2, t2)
	n3, t3 := palindrome.TimeNextGenerator(min, max)
	fmt.Printf("NextGenerator %d to %d: count %d: time %v\n", min, max, n3, t3)
	n4, t4 := palindrome.TimePrevGenerator(min, max)
	fmt.Printf("PrevGenerator %d to %d: count %d: time %v\n", min, max, n4, t4)
}
