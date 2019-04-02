package main

import (
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	// repetitions is the number of times to run the experiment.
	repetitions = 5

	// fileLength is the number of values to put into the random number file.
	fileLength = 100000

	// randomFile is the file to generate and input into each application.
	randomFile = "./randomFile.txt"

	// sortedFile is the file location that the application outputs.
	sortedFile = "./sortedFile.txt"

	// resultFile is the result file to write the duration of the application to.
	resultFile = "./results.txt"
)

// applications is the list of all compiled sort algorithms and languages.
// The result file will be written in the order of this list.
var applications = []string{}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

}

func randomizeApplicationOrder() ([]string, []int) {
	length := len(applications)
	appCopy := make([]string, length)
	copy(appCopy, applications)
	order := make([]int, length)
	for i := 0; i < length; i++ {
		order[i] = i
	}
	rand.Shuffle(length, func(i, j int) {
		appCopy[i], appCopy[j] = appCopy[j], appCopy[i]
		order[i], order[j] = order[j], order[i]
	})
	return appCopy, order
}

func createRandomFile() {
	values := make([]string, fileLength)
	for i := 0; i < fileLength; i++ {
		values[i] = strconv.Itoa(rand.Int())
	}
	data := strings.Join(values, "\n")
	if err := ioutil.WriteFile(randomFile, []byte(data), 0644); err != nil {
		panic(err)
	}
}
