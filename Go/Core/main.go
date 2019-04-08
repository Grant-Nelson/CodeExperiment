package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"sort"
	"strconv"
)

var (
	// inputFile is the file to input with random values from.
	inputFile = path.Join(`..`, `..`, `randomFile.txt`)

	// outputFile is the file to output the sorted values to.
	outputFile = path.Join(`..`, `..`, `sortedFile.txt`)
)

// main is the entry point for a core sort in Go.
func main() {
	data := readFile()
	length := len(data)
	if length <= 0 {
		panic(errors.New("Failed to read input file"))
	}

	sort.Ints(data)

	writeFile(data)
	os.Exit(0)
}

// readFile reads all the values from the input file.
func readFile() []int {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	randomData := []int{}
	reader := bufio.NewReader(file)
	for {
		data, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		value, err := strconv.Atoi(string(data))
		if err != nil {
			panic(err)
		}

		randomData = append(randomData, value)
	}

	return randomData
}

// writeFile writes the values to the output file.
func writeFile(data []int) {
	file, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for _, value := range data {
		file.WriteString(fmt.Sprintf("%d\n", value))
	}
}
