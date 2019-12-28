package main

import (
	"fmt"
)

func main() {
	wire1 := NewWire("input1.txt")
	wire2 := NewWire("input2.txt")
	crossWire := CrossWire{WireA: wire1, WireB: wire2}

	// intersection := crossWire.IntersectionClosestToCentral()

	// fmt.Println(intersection.DistanceFrom(centralPoint))
	fmt.Println(crossWire.FewestIntersectionStepsTotal())
}
