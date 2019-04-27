# Code Experiment

:seedling: **Analysing results, I will push them here when done, please come back later.**

-------

This project is for STAT-441 Exprimental Design at
Montana State University - Bozeman during Spring 2019
as part of my PhD program.

## Running Experiment

1. To change which treatments are built and run, modify
   the experiments [`main.go`](./experiment/main.go) file
2. Build experiment executable: `go build .\experiment\`
3. Build the treatments: `.\experiment.exe -build`
4. Run the experiment: `.\experiment.exe -run`
5. Results are collected in `results.txt` by default
6. Use `.\experiment.exe -h` to see more configurations

## Installing the Languages

This code was originally written for Windows 10 with
[Visual Studio Code](https://code.visualstudio.com/).

Many of the treatments and the experiment executable
should work on MacOS and Linux too but they have
only been tested on Windows.

### Install Instructions

- [C#](https://code.visualstudio.com/docs/languages/csharp#_installing-c35-support)
- [C++](https://code.visualstudio.com/docs/languages/cpp)
- [Go](https://golang.org/doc/install)
- [Java](https://www.azul.com/products/zulu-and-zulu-enterprise/download-openjdk-11)
- [Python](https://www.python.org/downloads/)
