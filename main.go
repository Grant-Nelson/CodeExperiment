package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
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
var applications = []string{
	"./Go/Binarytree/Binarytree.exe",
	"./Go/Mergesort/Mergesort.exe",
	"./Go/QuickSort/QuickSort.exe",
}

// main is the entry point for the experiment.
func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	resultF, err := os.Create(resultFile)
	if err != nil {
		panic(err)
	}
	defer resultF.Close()

	for i := 0; i < repetitions; i++ {
		runReplicate(resultF)
	}
}

// createRandomFile creates a new random file which can be used as input for a replication.
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

// randomizeApplicationOrder gets the randomized order of the applications and the paired order.
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

// checkOutputFile checks that the output file has been sorted.
func checkOutputFile(appName string) {
	file, err := os.Open(sortedFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	previousValue := -1
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

		if value < previousValue {
			panic(fmt.Errorf("%s did not properly sort, %d < %d", appName, value, previousValue))
		}
		previousValue = value
	}
	if previousValue < 0 {
		panic(fmt.Errorf("%s produced an empty file", appName))
	}
}

// runReplicate runs the set of the applications on a single input file.
func runReplicate(resultF *os.File) {
	createRandomFile()
	apps, order := randomizeApplicationOrder()
	results := make([]string, len(apps))

	for i, appName := range apps {
		cmd := exec.Command(appName)
		start := time.Now()

		if err := cmd.Run(); err != nil {
			panic(fmt.Sprintf("%s failed: %v", appName, err))
		}

		secs := time.Since(start).Seconds()
		checkOutputFile(appName)
		results[order[i]] = fmt.Sprintf("%.4f", secs)
	}

	resultF.WriteString(strings.Join(results, " ") + "\n")
	resultF.Sync()
}
