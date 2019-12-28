package main

import "math"

type Point [2]int

func (p *Point) DistanceFrom(point Point) int {
	x := math.Abs(float64(p[0]) - float64(point[0]))
	y := math.Abs(float64(p[1]) - float64(point[1]))

	return int(x + y)
}
