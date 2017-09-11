package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	dat, err := ioutil.ReadFile("data/simple.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Print(string(dat))
}
