package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func fetchInput() (modules Modules) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		i, err := strconv.Atoi(line)
		m := Module{i}
		modules = append(modules, m)

		if err != nil {
			log.Fatal(err)
		}
	}

	return
}

func main() {
	modules := fetchInput()

	fmt.Println(modules.totalFuel())
}
