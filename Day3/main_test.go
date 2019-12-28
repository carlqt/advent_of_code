package main

import "testing"

func TestPointIntersections(t *testing.T) {
	x := [][2]int{
		[2]int{4, 5},
		[2]int{1, 2},
		[2]int{7, 8},
	}

	y := [][2]int{
		[2]int{2, 1},
		[2]int{8, 7},
		[2]int{4, 5},
	}

	expected := [2]int{4, 5}

	intersections := PointIntersections(x, y)

	if intersections[0] != expected {
		t.Errorf("Intersections was incorrect, got: %v, want: %v", intersections[0], expected)
	}
}
