package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	out, err := exec.Command("sh", "-c", "ps -f").Output()
	if err != nil {
		log.Fatalf("Can't execute command: %s", err)
	}
	fmt.Printf("%s\n", out)
}
