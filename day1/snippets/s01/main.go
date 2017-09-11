package main

import (
	"fmt"
	"os"
)

var greeting string

func init() {
	glang := os.Getenv("LG_LANG")
	switch glang {
	case "de":
		greeting = "Hallo und willkommen!"
	default:
		greeting = "Let's Go!"
	}
}

func main() {
	fmt.Println(greeting)
	fmt.Println("👋")
}
