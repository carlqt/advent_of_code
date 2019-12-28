package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Wire []Point

func NewWire(path string) Wire {
	var wire Wire

	// Convert file input to a CSV of strings
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}

	// Read each value and try to draw the points
	// instructions := string(b)

	for _, instruction := range instructions(b) {
		wire.draw(instruction)
	}

	return wire
}

func instructions(input []byte) []string {
	stringInput := string(input)
	trimmedString := strings.TrimSuffix(stringInput, "\n")
	return strings.Split(trimmedString, ",")
}

func (w *Wire) CurrentPoint() [2]int {
	if w.Length() == 0 {
		return Point{0, 0}
	}

	return (*w)[w.Length()-1]
}

func (w *Wire) Length() int {
	return len(*w)
}

func (w *Wire) draw(instruction string) {
	direction := string(instruction[0])
	displacement, _ := strconv.Atoi(instruction[1:])

	switch direction {
	case "U":
		w.Up(displacement)
	case "D":
		w.Down(displacement)
	case "L":
		w.Left(displacement)
	default:
		w.Right(displacement)
	}
}

func (w *Wire) Up(steps int) {
	point := w.CurrentPoint()
	y := point[1]
	x := point[0]

	for ; y <= point[1]+steps; y++ {
		if point[1] == y {
			continue
		}

		*w = append(*w, Point{x, y})
	}
}

func (w *Wire) Down(steps int) {
	point := w.CurrentPoint()
	y := point[1]
	x := point[0]

	for ; y >= point[1]-steps; y-- {
		if point[1] == y {
			continue
		}

		*w = append(*w, Point{x, y})
	}
}

func (w *Wire) Left(steps int) {
	point := w.CurrentPoint()
	y := point[1]
	x := point[0]

	for ; x >= point[0]-steps; x-- {
		if point[0] == x {
			continue
		}

		*w = append(*w, Point{x, y})
	}
}

func (w *Wire) Right(steps int) {
	point := w.CurrentPoint()
	y := point[1]
	x := point[0]

	for ; x <= point[0]+steps; x++ {
		if point[0] == x {
			continue
		}

		*w = append(*w, Point{x, y})
	}
}

func (w *Wire) NumberOfSteps(point Point) (int, error) {
	for index, wirePoint := range *w {
		if wirePoint == point {
			return index + 1, nil
		}
	}

	return 0, errors.New("Point not found in wire")
}
