package main

import "testing"

func TestIntersections(t *testing.T) {
	x := []Point{
		Point{4, 5},
		Point{1, 2},
		Point{7, 8},
	}

	y := []Point{
		Point{2, 1},
		Point{8, 7},
		Point{4, 5},
	}

	expected := Point{4, 5}
	crossWire := CrossWire{WireA: x, WireB: y}

	intersections := crossWire.Intersections()

	if intersections[0] != expected {
		t.Errorf("Intersections was incorrect, got: %v, want: %v", intersections[0], expected)
	}
}

func TestFewestIntersectionSteps(t *testing.T) {
	wireA := Wire{
		Point{1, 0},
		Point{2, 0},
		Point{3, 0},
		Point{4, 0},
		Point{5, 0},
		Point{6, 0},
		Point{7, 0},
		Point{8, 0},
		Point{8, 1},
		Point{8, 2},
		Point{8, 3},
		Point{8, 4},
		Point{8, 5},
		Point{7, 5},
		Point{6, 5},
		Point{5, 5},
		Point{4, 5},
		Point{3, 5},
		Point{3, 4},
		Point{3, 3},
		Point{3, 3},
	}

	wireB := Wire{
		Point{0, 1},
		Point{0, 2},
		Point{0, 3},
		Point{0, 4},
		Point{0, 5},
		Point{0, 6},
		Point{0, 7},

		Point{1, 7},
		Point{2, 7},
		Point{3, 7},
		Point{4, 7},
		Point{5, 7},
		Point{6, 7},

		Point{6, 6},
		Point{6, 5},
		Point{6, 4},
		Point{6, 3},

		Point{5, 3},
		Point{4, 3},
		Point{3, 3},
		Point{2, 3},
	}

	crossWire := CrossWire{WireA: wireA, WireB: wireB}
	actual := crossWire.FewestIntersectionStepsTotal()
	expected := 30

	if actual != expected {
		t.Errorf("Steps total was incorrect, got: %d, want: %d", actual, expected)
	}
}
