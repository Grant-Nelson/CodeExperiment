package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
)

var (
	// inputFile is the file to input with random values from.
	inputFile = path.Join(`..`, `..`, `randomFile.txt`)

	// outputFile is the file to output the sorted values to.
	outputFile = path.Join(`..`, `..`, `sortedFile.txt`)
)

// quicksort performs a quick sort in the low inclusive and high inclusive range.
func quicksort(data []int, low, high int) {
	if low < high {
		p := low
		for j, pivot := low, data[high]; j < high; j++ {
			if data[j] < pivot {
				data[p], data[j] = data[j], data[p]
				p++
			}
		}
		data[p], data[high] = data[high], data[p]

		quicksort(data, low, p-1)
		quicksort(data, p+1, high)
	}
}

// sort will sort the given data.
func sort(data []int) {
	quicksort(data, 0, len(data)-1)
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

// main is the entry point for a quicksort in Go.
func main() {
	data := readFile()
	sort(data)
	writeFile(data)
	os.Exit(0)
}
