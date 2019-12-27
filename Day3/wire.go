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
	// fmt.Println(instructions)

	for _, instruction := range instructions(b) {
		wire.draw(instruction)
	}

	return wire
}

func instructions(input []byte) []string {
	stringInput := string(input)
	return strings.Split(stringInput, ",")
}

func (w *Wire) draw(instruction string) {
	direction := string(instruction[0])
	displacement, _ := strconv.Atoi(instruction[1:])
	point := w.CurrentPoint()
	x := point[0]
	y := point[1]

	switch direction {
	case "U":
		y = y + displacement
	case "D":
		y = y - displacement
	case "L":
		x = x - displacement
	default:
		x = x + displacement
	}

	w.Points = append(w.Points, [2]int{x, y})
}

func (w Wire) CurrentPoint() [2]int {
	if len(w.Points) == 0 {
		return [2]int{0, 0}
	}

	return w.Points[len(w.Points)-1]
}
