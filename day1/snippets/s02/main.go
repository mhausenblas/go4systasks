package main

import (
	"fmt"
)

func main() {
	cities := []string{"Dublin", "Berlin", "Rome"}
	for _, city := range cities {
		fmt.Println(city)
	}
	fmt.Println("---")
	cities[1] = "Munich"
	cities = append(cities, "Oslo")
	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}
}
