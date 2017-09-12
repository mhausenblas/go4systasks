package main

import (
	"fmt"
	"os"
)

func main() {
	// BEGIN OMIT
	f, err := os.Create("data/out.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	n, err := f.WriteString("Some text we're writing\nAnd another line")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("I wrote %d bytes\n", n)
	f.Sync()
	// END OMIT
}
