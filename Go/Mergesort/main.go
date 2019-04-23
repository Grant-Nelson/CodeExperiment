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

// sort will sort the given data.
func sort(data []int) {
	length := len(data)
	sortedData := make([]int, length)
	copy(sortedData, data)
	split(data, sortedData, 0, length)
	copy(data, sortedData)
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

// main is the entry point for a merge sort in Go.
func main() {
	data := readFile()
	sort(data)
	writeFile(data)
	os.Exit(0)
}
