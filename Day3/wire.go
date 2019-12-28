package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Wire struct {
	Points [][2]int
}

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

		w.Points = append(w.Points, [2]int{x, y})
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

		w.Points = append(w.Points, [2]int{x, y})
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

		w.Points = append(w.Points, [2]int{x, y})
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

		w.Points = append(w.Points, [2]int{x, y})
	}
}

func (w Wire) CurrentPoint() [2]int {
	if len(w.Points) == 0 {
		return [2]int{0, 0}
	}

	return w.Points[len(w.Points)-1]
}
