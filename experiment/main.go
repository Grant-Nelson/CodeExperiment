package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	// repetitions is the number of times to run the experiment.
	repetitions = 5

	// fileLength is the number of values to put into the random number file.
	fileLength = 100000

	// randomFile is the file to generate and input into each application.
	randomFile = `randomFile.txt`

	// sortedFile is the file location that the application outputs.
	sortedFile = `sortedFile.txt`

	// resultFile is the result file to write the duration of the application to.
	resultFile = `results.txt`

	// resultPrecision is the precision for the result values.
	resultPrecision = 5

	// resultSeparators is the characters to put between results.
	resultSeparators = `, `
)

// treatments is the list of all sort algorithms and languages.
// The result file will be written in the order of this list.
var treatments = []*treatment{

	goTreatment(`Binarytree`),
	goTreatment(`Mergesort`),
	goTreatment(`Quicksort`),

	javaTreatment(`Binarytree`),
	javaTreatment(`Mergesort`),
	javaTreatment(`Quicksort`),

	pythonTreatment(`Binarytree`),
	pythonTreatment(`Mergesort`),
	pythonTreatment(`Quicksort`),

	cSharpTreatment(`Binarytree`),
	cSharpTreatment(`Mergesort`),
	cSharpTreatment(`Quicksort`),
}

// main is the entry point for the experiment.
func main() {
	build := false
	run := false
	flag.BoolVar(&build, "build", false, "Set to build the treatment")
	flag.BoolVar(&run, "run", false, "Set to run the experiment")
	flag.Parse()

	if !build && !run {
		fmt.Println(`Must set "build" and/or "run" flag`)
		flag.PrintDefaults()
		os.Exit(1)
	}

	if build {
		for _, trmt := range treatments {
			trmt.build()
		}
	}

	if run {
		runExperiment()
	}

	os.Exit(0)
}
