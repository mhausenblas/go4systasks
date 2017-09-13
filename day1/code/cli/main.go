package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// BEGIN OMIT
func main() {
	var target string
	flag.StringVar(&target, "target", "", "The name of the person to praise.")
	repeat := flag.Int("repeat", 5, "Number of times to repeat the praise.")

	flag.Parse()
	if target == "" {
		log.Fatal("Can't praise an unspecified person")
	}
	for i := 0; i < *repeat; i++ {
		fmt.Printf("%s you are the best!\n", target)
	}
	fmt.Println("All arguments: ", os.Args)
}

// END OMIT
