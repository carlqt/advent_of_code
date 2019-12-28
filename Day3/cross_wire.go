package main

type CrossWire struct {
	WireA []Point
	WireB []Point
}

// func (cw CrossWire) Intersections
func (cw CrossWire) IntersectionClosestToCentral() Point {
	var closestPoint Point
	centralPoint := Point{0, 0}
	intersections := cw.Intersections()

	for index, point := range intersections {
		distance := point.DistanceFrom(centralPoint)

		if index == 0 {
			closestPoint = point
			continue
		}

		if closestPoint.DistanceFrom(centralPoint) > distance {
			closestPoint = point
		}
	}

	return closestPoint
}

func (cw *CrossWire) Intersections() []Point {
	// var dict map[[2]int]bool
	dict := make(map[Point]bool)

	// make a map of point1
	// initialize to false
	for _, point := range cw.WireA {
		dict[point] = false
	}

	// compare point2
	// if point is found, set value to true
	for _, point := range cw.WireB {
		_, ok := dict[point]

		if ok {
			dict[point] = true
		}
	}

	return intersectedPoints(dict)
}

func intersectedPoints(dict map[Point]bool) []Point {
	var result []Point

	for key, value := range dict {
		if value == true {
			result = append(result, key)
		}
	}

	return result
}
