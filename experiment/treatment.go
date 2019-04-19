package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"time"
)

const (
	// pathSep is the seperator between paths as a string.
	pathSep = string(os.PathSeparator)
)

// treatment defines how to run an experiment.
type treatment struct {

	// language is the programming language of the treatment.
	language string

	// algorithm is the implemented algorithm.
	algorithm string

	// index is the automatically defined zero-based order number
	// of this treatment in the list of treatments.
	index int

	// buildCmd is the executable name to build treatment with.
	// If an empty or nil, no build is performed.
	buildCmd []string

	// runCmd is the executable name to run treatment with.
	runCmd []string
}

// cppTreatment creates a C++ treatment for the given algorithm.
// This build is designed for Windows, must change for Linux and MacOS.
func cppTreatment(algorithm string) *treatment {
	cppExe := `C:\mingw-w64\i686-8.1.0-posix-dwarf-rt_v6-rev0\mingw32\bin\g++.exe`
	return &treatment{
		language:  `C++`,
		algorithm: algorithm,
		index:     -1,
		buildCmd:  []string{cppExe, `main.cpp`, `-o`, `main`},
		runCmd:    []string{fmt.Sprint(`.`, pathSep, `main.exe`)},
	}
}

// cSharpTreatment creates a C# treatment for the given algorithm.
// This build is designed for Windows, must change for Linux and MacOS.
func cSharpTreatment(algorithm string) *treatment {
	cscExe := `C:\Windows\Microsoft.NET\Framework\v4.0.30319\csc.exe`
	return &treatment{
		language:  `C#`,
		algorithm: algorithm,
		index:     -1,
		buildCmd:  []string{cscExe, `/t:exe`, `/out:main.exe`, `main.cs`},
		runCmd:    []string{fmt.Sprint(`.`, pathSep, `main.exe`)},
	}
}

// goTreatment creates a Go treatment for the given algorithm.
func goTreatment(algorithm string) *treatment {
	return &treatment{
		language:  `Go`,
		algorithm: algorithm,
		index:     -1,
		buildCmd:  []string{`go`, `build`},
		runCmd:    []string{fmt.Sprint(`.`, pathSep, algorithm, `.exe`)},
	}
}

// javaTreatment creates a Java treatment for the given algorithm.
func javaTreatment(algorithm string) *treatment {
	return &treatment{
		language:  `Java`,
		algorithm: algorithm,
		index:     -1,
		buildCmd:  []string{`javac`, algorithm + `.java`},
		runCmd:    []string{`java`, algorithm},
	}
}

// pythonTreatment creates a Python treatment for the given algorithm.
func pythonTreatment(algorithm string) *treatment {
	return &treatment{
		language:  `Python`,
		algorithm: algorithm,
		index:     -1,
		buildCmd:  nil, // nothing needs to be done for python
		runCmd:    []string{`python`, `main.py`},
	}
}

// String gets the display name of the treatment.
func (trmt *treatment) String() string {
	return trmt.language + "-" + trmt.algorithm
}

// Path gets the path to the treatment code.
func (trmt *treatment) Path() string {
	return path.Join(trmt.language, trmt.algorithm)
}

// build runs the build command for preparing the treatment.
func (trmt *treatment) Build() {
	fmt.Printf("building %s\n", trmt.String())
	if len(trmt.buildCmd) > 0 {
		cmd := exec.Command(trmt.buildCmd[0], trmt.buildCmd[1:]...)
		cmd.Dir = trmt.Path()
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			panic(fmt.Sprintf("%s failed to build: %v", trmt.String(), err))
		}
	}
}

// run executes the treatment and returns the time in seconds it took to execute.
func (trmt *treatment) Run() float64 {
	cmd := exec.Command(trmt.runCmd[0], trmt.runCmd[1:]...)
	cmd.Dir = trmt.Path()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	fmt.Printf("  running %s...", trmt.String())

	start := time.Now()
	if err := cmd.Run(); err != nil {
		panic(fmt.Sprintf("%s failed to run: %v", trmt.String(), err))
	}
	dur := time.Since(start)

	fmt.Printf("%s\n", dur.String())
	return dur.Seconds()
}
