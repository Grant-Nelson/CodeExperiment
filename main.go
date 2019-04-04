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
	randomFile = `.\randomFile.txt`

	// sortedFile is the file location that the application outputs.
	sortedFile = `.\sortedFile.txt`

	// resultFile is the result file to write the duration of the application to.
	resultFile = `.\results.txt`

	// resultPrecision is the precision for the result values.
	resultPrecision = 5
)

// applications is the list of all compiled sort algorithms and languages.
// The result file will be written in the order of this list.
var applications = [][]string{
	// Name, Path, Executable, Args...
	{`Go-Binarytree`, `.\Go\Binarytree\`, `.\Binarytree.exe`},
	{`Go-Mergesort`, `.\Go\Mergesort\`, `.\Mergesort.exe`},
	{`Go-Quicksort`, `.\Go\QuickSort\`, `.\QuickSort.exe`},
	{`Java-Binarytree`, `.\Java\Binarytree`, `java`, `Binarytree`},
	{`Java-Mergesort`, `.\Java\Mergesort`, `java`, `Mergesort`},
	{`Java-Quicksort`, `.\Java\Quicksort`, `java`, `Quicksort`},
}

// main is the entry point for the experiment.
func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	resultF, err := os.Create(resultFile)
	if err != nil {
		panic(err)
	}
	defer resultF.Close()

	for i := 1; i <= repetitions; i++ {
		fmt.Printf("replicate %d\n", i)
		runReplicate(resultF)
	}
}

// createRandomFile creates a new random file which can be used as input for a replication.
func createRandomFile() {
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
func randomizeApplicationOrder() ([][]string, []int) {
	length := len(applications)

	appCopy := make([][]string, length)
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
			panic(fmt.Errorf("%s did not properly sort at line %d, %d < %d", appName, count, value, previousValue))
		}
		previousValue = value
	}
	if count != fileLength {
		panic(fmt.Errorf("%s produced a file with only %d values instead of %d", appName, count, fileLength))
	}
}

// deleteSortedFile deletes the result file.
func deleteSortedFile() {
	if err := os.Remove(sortedFile); err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}
	}
}

// runReplicate runs the set of the applications on a single input file.
func runReplicate(resultF *os.File) {
	createRandomFile()
	deleteSortedFile()
	apps, order := randomizeApplicationOrder()
	results := make([]string, len(apps))
	for i, appCmd := range apps {
		name := appCmd[0]
		cmd := exec.Command(appCmd[2], appCmd[3:]...)
		cmd.Dir = appCmd[1]
		fmt.Printf("  running %s\n", name)

		start := time.Now()
		if err := cmd.Run(); err != nil {
			panic(fmt.Sprintf("%s failed: %v", name, err))
		}
		secs := time.Since(start).Seconds()

		checkOutputFile(name)
		deleteSortedFile()
		results[order[i]] = fmt.Sprintf("%.*f", resultPrecision, secs)
	}

	resultF.WriteString(strings.Join(results, " ") + "\n")
	resultF.Sync()
}
