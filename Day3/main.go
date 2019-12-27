package main

import (
	"math"
)

func main() {
	// wire1 := NewWire("input1.txt")
	// wire2 := NewWire("input2.txt")

	// fmt.Println(distance([2]int{3, 3}, [2]int{0, 0}))
	// read inputs
	// create a function to get all intersections
	// create a function to calculate the distance of the points of intersections
}

func distance(point1 [2]int, point2 [2]int) int {
	x := math.Abs(float64(point1[0]) - float64(point2[0]))
	y := math.Abs(float64(point1[1]) - float64(point2[1]))

	return int(x + y)
}

// func pointIntersections() [][2]int {
// }
