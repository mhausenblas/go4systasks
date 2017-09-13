package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// BEGIN OMIT
	now := time.Now()
	fmt.Printf("Now (Unix timestamp): %v\n", now.Unix())
	fmt.Printf("Now (RFC 3339): %s\n", now.Format(time.RFC3339))
	fmt.Printf("Now (custom): %s\n", now.Format("Jan _2 2006, 15:04"))

	sometime, err := time.Parse(time.RFC3339, "2017-04-01T14:17:09+00:00")
	if err != nil {
		log.Printf("Can't parse time: %v", err)
	}
	fmt.Printf("Sometime: %s\n", sometime)

	diff := now.Sub(sometime)
	fmt.Printf("Difference: %v or ca. %3.0f days\n", diff, diff.Hours()/24)
	// END OMIT
}
