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

// node is the binary tree node.
type node struct {

	// value is the value for this node.
	value int

	// left is the value less than this value.
	left *node

	// right is the value greater than or equal to this value.
	right *node
}

// newNode creates a new binary tree node.
func newNode(value int) *node {
	return &node{
		value: value,
		left:  nil,
		right: nil,
	}
}

// main is the entry point for a binary tree sort in Go.
func main() {
	data := readFile()
	length := len(data)
	if length <= 0 {
		panic(errors.New("Failed to read input file"))
	}

	root := newNode(data[0])
	for i := 1; i < length; i++ {
		insertValue(data[i], root)
	}

	outputValues(0, root, data)

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

// insertValue inserts a value into the tree recursively.
func insertValue(value int, n *node) {
	if n.value > value {
		if n.left != nil {
			insertValue(value, n.left)
		} else {
			n.left = newNode(value)
		}
	} else {
		if n.right != nil {
			insertValue(value, n.right)
		} else {
			n.right = newNode(value)
		}
	}
}

// outputValues recursively gets all the values from the binary tree.
func outputValues(index int, n *node, data []int) int {
	if n.left != nil {
		index = outputValues(index, n.left, data)
	}

	data[index] = n.value
	index++

	if n.right != nil {
		index = outputValues(index, n.right, data)
	}
	return index
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
