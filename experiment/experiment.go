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

const (
	// randomFile is the file to generate and input into each application.
	randomFile = `randomFile.txt`

	// sortedFile is the file location that the application outputs.
	sortedFile = `sortedFile.txt`

	// resultPrecision is the precision for the result values.
	resultPrecision = 5
)

// experiment is an object for running an experiment.
type experiment struct {

	// order is a monotonically increasing treatment order number.
	order int

	// startTime is the time the experiment started at.
	startTime time.Time

	// repetitions is the number of times to run the experiment.
	repetitions int

	// fileLength is the number of values to put into the random number file.
	fileLength int

	// resultFile is the result file to write the duration of the application to.
	resultFile string

	// fixFile indicates if (true) only one random file should be used for all repetitions,
	// otherwise (false) a new random file is created per repetition.
	fixFile bool
}

// newExperiment constructs a new experiment.
func newExperiment(repetitions, fileLength int, resultFile string, fixFile bool) *experiment {
	return &experiment{
		order:       1,
		startTime:   time.Now(),
		repetitions: repetitions,
		fileLength:  fileLength,
		resultFile:  resultFile,
		fixFile:     fixFile,
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
	resultF, err := os.Create(e.resultFile)
	if err != nil {
		panic(err)
	}
	defer resultF.Close()
	resultF.WriteString("file length: " + strconv.Itoa(e.fileLength) + "\n")
	resultF.WriteString("fixed file:  " + strconv.FormatBool(e.fixFile) + "\n")
	resultF.WriteString("order replicate index language algorithm seconds\n")

	// Run all the repetitions of the experiment.
	e.order = 1
	e.startTime = time.Now()
	for i := 1; i <= e.repetitions; i++ {
		e.runReplicate(i, resultF)
	}
}

// runReplicate runs the set of the applications on a single input file.
func (e *experiment) runReplicate(replicate int, resultF *os.File) {
	fmt.Printf("replicate %d of %d\n", replicate, e.repetitions)
	repStartTime := time.Now()

	if (!e.fixFile) || (replicate == 1) {
		e.createRandomFile()
	}

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
	repDur := time.Since(repStartTime)
	remaining := time.Duration(float64(e.repetitions-replicate) * float64(time.Since(e.startTime)) / float64(replicate))
	fmt.Printf("  took " + repDur.String() + ", about " + remaining.String() + " remaining\n")
}

// createRandomFile creates a new random file which can be used as input for a replication.
func (e *experiment) createRandomFile() {
	values := make([]string, e.fileLength)
	for i := 0; i < e.fileLength; i++ {
		// Use 0 to 2^31 to work with all 32-bit base languages
		values[i] = strconv.Itoa(int(rand.Int31()))
	}

	data := strings.Join(values, "\n")
	if err := ioutil.WriteFile(randomFile, []byte(data), 0644); err != nil {
		panic(err)
	}

	fmt.Printf("  input file generated with %d random numbers\n", e.fileLength)
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
	if count != e.fileLength {
		panic(fmt.Errorf("%s produced a file with only %d values instead of %d", trmt.String(), count, e.fileLength))
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
