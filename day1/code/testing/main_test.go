package main

import (
	"testing"
)

// BEGIN OMIT
func TestSum(t *testing.T) {
	in0 := 1
	in1 := 1
	want := 2
	got := sum(in0, in1)
	if got != want {
		t.Errorf("sum(%d, %d) == %d, want %d", in0, in1, got, want)
	}
}

// END OMIT
