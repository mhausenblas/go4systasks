package main

import (
	"fmt"
	"log"
	"strconv"
)

// BEGIN OMIT
func main() {
	fmt.Printf("What's your preferred number?\n> ")
	var input string
	fmt.Scanln(&input)
	fmt.Printf("You said: %v\n", input)
	num, err := strconv.Atoi(input)
	if err != nil {
		log.Fatalf("Can't parse '%v' as a number: %s\n", input, err)
	}
	if num == 42 {
		fmt.Println("I knew it, good choice!")
	}
}

//END OMIT
