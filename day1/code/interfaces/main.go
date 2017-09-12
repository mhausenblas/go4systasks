package main

import (
	"fmt"
)

// BEGIN OMIT
type num int

// Formatter returns a formatted version of a value as a string.
type Formatter interface {
	Format() string
}

func (n num) Format() string {
	return fmt.Sprintf("%d", n)
}

func printf(f Formatter) {
	fmt.Println(f.Format())
}

func main() {
	anumber := num(42)
	printf(anumber)
}

// END OMIT
