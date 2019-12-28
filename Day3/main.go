package main

import (
	"fmt"
	"math"
)

func main() {
	wire1 := NewWire("input1.txt")
	wire2 := NewWire("input2.txt")

	intersections := PointIntersections(wire1.Points, wire2.Points)

	distance := shortestDistance(intersections)
	fmt.Println(distance)

	// fmt.Println(distance([2]int{3, 3}, [2]int{0, 0}))
	// read inputs
	// create a function to get all intersections
	// create a function to calculate the distance of the points of intersections
}

func calculateDistance(point1 [2]int, point2 [2]int) int {
	x := math.Abs(float64(point1[0]) - float64(point2[0]))
	y := math.Abs(float64(point1[1]) - float64(point2[1]))

	return int(x + y)
}

func shortestDistance(points [][2]int) int {
	var distance int
	centralPoint := [2]int{0, 0}

	for index, point := range points {
		if point == centralPoint {
			continue
		}

		pointDistance := calculateDistance(point, centralPoint)

		if index == 0 {
			distance = pointDistance
		}

		if distance > pointDistance {
			distance = pointDistance
		}
	}

	return distance
}

func PointIntersections(point1 [][2]int, point2 [][2]int) [][2]int {
	// var dict map[[2]int]bool
	dict := make(map[[2]int]bool)

	// make a map of point1
	// initialize to false
	for _, point := range point1 {
		dict[point] = false
	}

	// compare point2
	// if point is found, set value to true
	for _, point := range point2 {
		_, ok := dict[point]

		if ok {
			dict[point] = true
		}
	}

	// return the keys with value of true
	// return [][2]int{[2]int{4, 4}}
	return intersectedPoints(dict)
}

func intersectedPoints(dict map[[2]int]bool) [][2]int {
	var result [][2]int

	for key, value := range dict {
		if value == true {
			result = append(result, key)
		}
	}

	return result
}
