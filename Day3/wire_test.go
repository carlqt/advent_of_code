package main

import "testing"

func TestNumberOfSteps(t *testing.T) {
	wire := Wire{
		Point{4, 5},
		Point{1, 2},
		Point{7, 8},
	}

	actual, _ := wire.NumberOfSteps(Point{7, 8})
	expected := 3

	if actual != expected {
		t.Errorf("Result was incorrect, got: %d, want: %d", actual, expected)
	}

	// when there is an error
	_, err := wire.NumberOfSteps(Point{99, 99})

	if err == nil {
		t.Errorf("Result was incorrect, no error returned")
	}
}
