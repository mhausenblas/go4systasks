package main

import "fmt"

func something(this string, that *string) {
	fmt.Println(this, *that)
}

func someother(some []int) {
	for _, e := range some {
		fmt.Print(e)
	}
	fmt.Println("")
	some[1] = 100
}

func pointersome(ref *[]int) {
	(*ref)[2] = 42
}

func main() {

	that := "world"
	something("hello", &that)

	intlist := []int{1, 2, 3}

	pointersome(&intlist)
	someother(intlist)

}
