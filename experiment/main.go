package main

import (
	"flag"
	"fmt"
	"os"
)

// treatments is the list of all sort algorithms and languages.
// The result file will be written in the order of this list.
var treatments = []*treatment{

	cSharpTreatment(`Binarytree`),
	//cSharpTreatment(`Core`),
	cSharpTreatment(`Mergesort`),
	cSharpTreatment(`Quicksort`),

	cppTreatment(`Binarytree`),
	//cppTreatment(`Core`),
	cppTreatment(`Mergesort`),
	cppTreatment(`Quicksort`),

	goTreatment(`Binarytree`),
	//goTreatment(`Core`),
	goTreatment(`Mergesort`),
	goTreatment(`Quicksort`),

	javaTreatment(`Binarytree`),
	//javaTreatment(`Core`),
	javaTreatment(`Mergesort`),
	javaTreatment(`Quicksort`),

	pythonTreatment(`Binarytree`),
	//pythonTreatment(`Core`),
	pythonTreatment(`Mergesort`),
	pythonTreatment(`Quicksort`),
}

// main is the entry point for the experiment.
func main() {
	build := false
	flag.BoolVar(&build, "build", build,
		`Set to build the treatment.`)

	run := false
	flag.BoolVar(&run, "run", run,
		`Set to run the experiment.`)

	repetitions := 100
	flag.IntVar(&repetitions, "repetitions", repetitions,
		`The number of times to run the experiment.`)

	fileLength := 10000
	flag.IntVar(&fileLength, "fileLength", fileLength,
		`The number of values to put into the random number file.`)

	resultFile := `results.txt`
	flag.StringVar(&resultFile, "out", resultFile,
		`The result file to write the duration of the application to.`)

	fixFile := false
	flag.BoolVar(&fixFile, "fix", fixFile,
		`Indicates that only one random file should be used for all repetitions.`)

	flag.Parse()

	if !build && !run {
		fmt.Println(`Must set "build" and/or "run" flag`)
		flag.PrintDefaults()
		os.Exit(1)
	}

	if build {
		for _, trmt := range treatments {
			trmt.Build()
		}
	}

	if run {
		newExperiment(repetitions, fileLength, resultFile, fixFile).runExperiment()
	}

	os.Exit(0)
}
