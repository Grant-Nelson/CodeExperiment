package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

const (
	// inputFile is the file to input with random values from.
	inputFile = "../../randomFile.txt"

	// outputFile is the file to output the sorted values to.
	outputFile = "../../sortedFile.txt"
)

// main is the entry point for a quicksort in Go.
func main() {
	data := readFile()
	length := len(data)
	if length <= 0 {
		panic(errors.New("Failed to read a file"))
	}

	quicksort(data, 0, length-1)
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

// quicksort performs a quick sort in the low inclusive and high inclusive range.
func quicksort(data []int, low, high int) {
	if low < high {
		p := partition(data, low, high)
		quicksort(data, low, p-1)
		quicksort(data, p+1, high)
	}
}

// partition shifts values lower than a pivot and returns the pivot index.
func partition(data []int, low, high int) int {
	pivot := data[high]
	i := low
	for j := low; j < high; j++ {
		if data[j] < pivot {
			data[i], data[j] = data[j], data[i]
			i++
		}
	}
	data[i], data[high] = data[high], data[i]
	return i
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
