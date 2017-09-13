package main

import (
	"fmt"
)

// BEGIN OMIT
func main() {
	a := 40
	b := 2
	fmt.Printf("We found that: %d + %d = %d\n", a, b, sum(a, b))
}

func sum(a, b int) int {
	return a + b
}

// END OMIT
