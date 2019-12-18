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

	for i := 0; i < len(inputs); i = i + 4 {
		endSlice := i + 4
		command := inputs[i:endSlice]

		if command[0] == 99 {
			log.Printf("Exiting")
			break
		} else {
			execute(command, inputs)
		}
	}

	fmt.Println(inputs[0])
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
	log.Printf("Executing %v\n", command)

	var result int
	opcode := command[0]

	indexOfInput1 := command[1]
	indexOfInput2 := command[2]
	input1 := intCodeInput[indexOfInput1]
	input2 := intCodeInput[indexOfInput2]
	position := command[3]

	if opcode == 1 {
		result = input1 + input2
	} else if opcode == 2 {
		result = input1 * input2
	} else {
		log.Printf("Unknown opcode %d for command %v", opcode, command)
	}

	intCodeInput[position] = result
}
