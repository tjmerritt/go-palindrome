
// Copyright (C) 2018 Thomas J. Merrtt

package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/tjmerritt/go-palindrome"
)

var min int = 1
var max int = 1000

func main() {
	flag.IntVar(&min, "min", 1, "Minimum palindrome value")
	flag.IntVar(&max, "max", 1000, "Maximum palindrome value")
	flag.Parse()

	p := min
	
	if p > 9 {
		p = palindrome.Next(p - 1)
	}

	fmt.Printf("[\n")
	count := 0
	st := time.Now()

	for p <= max {
		count++
		fmt.Printf("    %d,\n", p)
		p = palindrome.Next(p)
	}

	t := time.Now().Sub(st)
	fmt.Printf("]\n")
	fmt.Printf("Range: %d %d\n", min, max)
	fmt.Printf("Palindromes: %d\n", count)
	fmt.Printf("Run time: %v\n", t)
	return
}
