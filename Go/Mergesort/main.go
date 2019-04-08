package main

import (
	"bufio"
	"errors"
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

// main is the entry point for a merge sort in Go.
func main() {
	randomData := readFile()
	length := len(randomData)
	if length <= 0 {
		panic(errors.New("Failed to read input file"))
	}

	sortedData := make([]int, length)
	copy(sortedData, randomData)
	split(randomData, sortedData, 0, length)

	writeFile(sortedData)
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

// split performs a top down merge sort by splitting the
// current level into 2 parts to sort, then merging the two parts.
// start is inclusive and stop is exclusive.
func split(a, b []int, start, stop int) {
	if stop-start < 2 {
		return
	}

	mid := (stop + start) / 2
	split(b, a, start, mid)
	split(b, a, mid, stop)

	for i, j, k := start, mid, start; k < stop; k++ {
		if (i < mid) && ((j >= stop) || (a[i] <= a[j])) {
			b[k] = a[i]
			i++
		} else {
			b[k] = a[j]
			j++
		}
	}
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
