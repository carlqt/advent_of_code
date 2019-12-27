package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

const EXPECTED_RESULT = 19690720

func main() {
	memory := intCodeInput("input.csv")

	bruteForceInput(memory)
}
func bruteForceInput(memory []int) {
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			newMemory := make([]int, len(memory))
			copy(newMemory, memory)

			newMemory[1] = x
			newMemory[2] = y

			address := execute(newMemory)

			if address == EXPECTED_RESULT {
				result := 100*newMemory[1] + newMemory[2]
				fmt.Println(result)
				return
			}
		}
	}
}

func execute(memory []int) int {
	for i := 0; i < len(memory); i = i + 4 {
		endSlice := i + 4
		instruction := NewInstruction(memory[i:endSlice], memory)

		if instruction.Halt() {
			break
		} else {
			instruction.Execute(memory)
		}
	}

	return memory[0]
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
