package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// BEGINC OMIT
	now := time.Now()
	fmt.Printf("Now (RFC 3339): %s\n", now.Format(time.RFC3339))

	sometime, err := time.Parse(time.RFC3339, "2017-01-01T21:09:09+00:00")
	if err != nil {
		log.Printf("Can't parse time: %v", err)
	}
	fmt.Printf("Sometime: %s\n", sometime)

	diff := now.Sub(sometime)
	fmt.Printf("Difference: %v or ca. %3.0f days\n", diff, diff.Hours()/24)
	// ENDC OMIT
}
