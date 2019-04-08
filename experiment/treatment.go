package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"time"
)

// treatment defines how to run an experiment.
type treatment struct {

	// name is the display name of the treatment.
	name string

	// path is the path to the treatment package.
	path string

	// buildCmd is the executable name to build treatment with.
	// If an empty or nil, no build is performed.
	buildCmd []string

	// runCmd is the executable name to run treatment with.
	runCmd []string

	// index is the automatically defined zero-based order number
	// of this treatment in the list of treatments.
	index int
}

// cSharpTreatment creates a C# treatment for the given name.
// This build is designed for Windows, must change for Linux and MacOS.
func cSharpTreatment(name string) *treatment {
	cscExe := `C:\Windows\Microsoft.NET\Framework\v4.0.30319\csc.exe`
	return &treatment{
		name:     `C#-` + name,
		path:     path.Join(`Csharp`, name),
		buildCmd: []string{cscExe, `/t:exe`, `/out:main.exe`, `main.cs`},
		runCmd:   []string{fmt.Sprint(`.`, string(os.PathSeparator), `main.exe`)},
	}
}

// goTreatment creates a Go treatment for the given name.
func goTreatment(name string) *treatment {
	return &treatment{
		name:     `Go-` + name,
		path:     path.Join(`Go`, name),
		buildCmd: []string{`go`, `build`},
		runCmd:   []string{fmt.Sprint(`.`, string(os.PathSeparator), name, `.exe`)},
	}
}

// javaTreatment creates a Java treatment for the given name.
func javaTreatment(name string) *treatment {
	return &treatment{
		name:     `Java-` + name,
		path:     path.Join(`Java`, name),
		buildCmd: []string{`javac`, name + `.java`},
		runCmd:   []string{`java`, name},
	}
}

// pythonTreatment creates a Python treatment for the given name.
func pythonTreatment(name string) *treatment {
	return &treatment{
		name:     `Python-` + name,
		path:     path.Join(`Python`, name),
		buildCmd: nil, // nothing needs to be done for python
		runCmd:   []string{`python`, `main.py`},
	}
}

// build runs the build command for preparing the treatment.
func (trmt *treatment) build() {
	fmt.Printf("building %s\n", trmt.name)
	if len(trmt.buildCmd) > 0 {
		cmd := exec.Command(trmt.buildCmd[0], trmt.buildCmd[1:]...)
		cmd.Dir = trmt.path
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			panic(fmt.Sprintf("%s failed to build: %v", trmt.name, err))
		}
	}
}

// run executes the treatment and returns the time in seconds it took to execute.
func (trmt *treatment) run() float64 {
	cmd := exec.Command(trmt.runCmd[0], trmt.runCmd[1:]...)
	cmd.Dir = trmt.path
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	fmt.Printf("  running %s...", trmt.name)

	start := time.Now()
	if err := cmd.Run(); err != nil {
		panic(fmt.Sprintf("%s failed to run: %v", trmt.name, err))
	}
	dur := time.Since(start)

	fmt.Printf("%s\n", dur.String())
	return dur.Seconds()
}
