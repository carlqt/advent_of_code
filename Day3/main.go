package main

import (
	"fmt"
)

func main() {
	wire1 := NewWire("input1.txt")
	wire2 := NewWire("input2.txt")
	centralPoint := Point{0, 0}
	crossWire := CrossWire{WireA: wire1, WireB: wire2}

	intersection := crossWire.IntersectionClosestToCentral()

	fmt.Println(intersection.DistanceFrom(centralPoint))
}
