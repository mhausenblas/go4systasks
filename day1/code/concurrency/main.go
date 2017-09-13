package main

import "fmt"

// BEGIN OMIT
func generator(c chan string) {
	var prev string
	for i := 0; i < cap(c); i++ {
		prev = fmt.Sprintf("%s%d", prev, i)
		c <- prev
	}
	close(c)
}

func main() {
	c := make(chan string, 10)
	go generator(c)
	for i := range c {
		fmt.Println(i)
	}
}

// END OMIT
