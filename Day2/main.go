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
	memory := intCodeInput("input.csv")

	for i := 0; i < len(memory); i = i + 4 {
		endSlice := i + 4
		instruction := NewInstruction(memory[i:endSlice], memory)

		if instruction.Halt() {
			log.Printf("Exiting")
			break
		} else {
			instruction.Execute(memory)
		}
	}

	fmt.Println(memory[0])
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
