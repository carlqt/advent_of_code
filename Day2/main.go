package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	inputs := intCodeInput("input.csv")

	fmt.Println(len(inputs))
}

func intCodeInput(url string) []int {
	var record []string

	file, err := os.Open(url)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		body, err := reader.Read()
		record = append(record, body...)

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}

	return formatToInt(record)
}

// Assume input is safe
func formatToInt(record []string) []int {
	var result []int

	for _, r := range record {
		num, _ := strconv.Atoi(r)
		result = append(result, num)
	}

	return result
}

func execute(command []int, intCodeInput []int) {
	var result int
	opcode := command[0]

	indexOfInput1 := command[1]
	indexOfInput2 := command[2]
	input1 := intCodeInput[indexOfInput1]
	input2 := intCodeInput[indexOfInput2]
	position := command[3]

	if opcode == 99 {
		return
	} else if opcode == 1 {
		result = input1 + input2
	} else if opcode == 2 {
		result = input1 * input2
	} else {
		log.Fatal("Unknown opcode")
	}

	intCodeInput[position] = result
}
