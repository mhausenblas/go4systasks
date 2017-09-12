package main

import (
	"fmt"
)

func main() {
	seniority := map[string]int{"Tom": 12, "Ann": 20, "Megan": 37, "Sid": 46}
	seniority["Laura"] = 29
	for name, age := range seniority {
		fmt.Printf("The age of %s is %d\n", name, age)
	}
}
