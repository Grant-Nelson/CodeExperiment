package main

import (
	"fmt"
	"os/exec"
)

// main is the entry point for building the experiment.
func main() {
	buildGo(`Binarytree`)
	buildGo(`Mergesort`)
	buildGo(`Quicksort`)

	buildJava(`Binarytree`)
	buildJava(`Mergesort`)
	buildJava(`Quicksort`)

	buildPython(`Binarytree`)
	buildPython(`Mergesort`)
	buildPython(`Quicksort`)
}

// buildGo builds the Go treatment with the given name.
func buildGo(name string) {
	cmd := exec.Command("go", "build")
	cmd.Dir = `.\Go\` + name
	if err := cmd.Run(); err != nil {
		panic(fmt.Sprintf("Failed to build GO at %s: %v", name, err))
	}
}

// buildJava builds the Java treatment with the given name.
func buildJava(name string) {
	cmd := exec.Command("javac", name+".java")
	cmd.Dir = `.\Java\` + name
	if err := cmd.Run(); err != nil {
		panic(fmt.Sprintf("Failed to build Java at %s: %v", name, err))
	}
}

// buildPython builds the Python treatment with the given name.
func buildPython(name string) {
	// nothing needs to be done for python
}
