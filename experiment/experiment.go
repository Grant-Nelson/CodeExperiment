package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// experiment is an object for running an experiment.
type experiment struct {

	// order is a monotonically increasing treatment order number.
	order int
}

// newExperiment constructs a new experiment.
func newExperiment() *experiment {
	return &experiment{
		order: 1,
	}
}

// runExperiment runs the full experiment.
func (e *experiment) runExperiment() {
	// Seed the random number generator with current nanoseconds.
	rand.Seed(time.Now().UTC().UnixNano())

	// Number all the selected treatments.
	names := make([]string, len(treatments))
	for i, trmt := range treatments {
		trmt.index = i
		names[i] = trmt.String()
	}

	// Create the results output file.
	resultF, err := os.Create(resultFile)
	if err != nil {
		panic(err)
	}
	defer resultF.Close()
	resultF.WriteString("order replicate index language algorithm seconds\n")

	// Run all the repetitions of the experiment.
	e.order = 1
	for i := 1; i <= repetitions; i++ {
		fmt.Printf("replicate %d of %d\n", i, repetitions)
		e.runReplicate(i, resultF)
	}
}

// runReplicate runs the set of the applications on a single input file.
func (e *experiment) runReplicate(replicate int, resultF *os.File) {
	e.createRandomFile()
	e.deleteSortedFile()
	trmts := e.randomizeApplicationOrder()
	for _, trmt := range trmts {

		// Run treatment
		secs := trmt.Run()

		// Check it rand it correctly
		e.checkOutputFile(trmt)
		e.deleteSortedFile()

		// Write results to the result file
		result := fmt.Sprintf("%d %d %d %s %s %.*f\n",
			e.order, replicate, trmt.index, trmt.language, trmt.algorithm, resultPrecision, secs)
		resultF.WriteString(result)
		e.order++
	}
	resultF.Sync()
}

// createRandomFile creates a new random file which can be used as input for a replication.
func (e *experiment) createRandomFile() {
	values := make([]string, fileLength)
	for i := 0; i < fileLength; i++ {
		// Use 0 to 2^31 to work with all 32-bit base languages
		values[i] = strconv.Itoa(int(rand.Int31()))
	}

	data := strings.Join(values, "\n")
	if err := ioutil.WriteFile(randomFile, []byte(data), 0644); err != nil {
		panic(err)
	}
}

// randomizeApplicationOrder gets the randomized order of the applications and the paired order.
func (e *experiment) randomizeApplicationOrder() []*treatment {
	length := len(treatments)
	trmtCopy := make([]*treatment, length)
	copy(trmtCopy, treatments)

	rand.Shuffle(length, func(i, j int) {
		trmtCopy[i], trmtCopy[j] = trmtCopy[j], trmtCopy[i]
	})

	return trmtCopy
}

// checkOutputFile checks that the output file has been sorted.
func (e *experiment) checkOutputFile(trmt *treatment) {
	file, err := os.Open(sortedFile)
	if err != nil {
		panic(fmt.Errorf("failed to read sorted output file from %s: %v", trmt.String(), err))
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	previousValue := -1
	count := 0
	for {
		data, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		count++
		value, err := strconv.Atoi(string(data))
		if err != nil {
			panic(err)
		}

		if value < previousValue {
			panic(fmt.Errorf("%s did not properly sort at line %d, %d < %d", trmt.String(), count, value, previousValue))
		}
		previousValue = value
	}
	if count != fileLength {
		panic(fmt.Errorf("%s produced a file with only %d values instead of %d", trmt.String(), count, fileLength))
	}
}

// deleteSortedFile deletes the result file.
func (e *experiment) deleteSortedFile() {
	if err := os.Remove(sortedFile); err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}
	}
}
