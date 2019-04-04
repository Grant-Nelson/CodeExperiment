# Code Experiment

:seedling: **This experiment is not ready yet! Come back later.**

- [Running Experiment](running-experiment)
- [Compiling Treatments](compiling-experiments)
  - [Go](go)
    - [Binarytree](go-binarytree)
    - [Mergesort](go-mergesort)
    - [Quicksort](go-quicksort)
  - [Java](java)
    - [Binarytree](java-binarytree)
    - [Mergesort](java-mergesort)
    - [Quicksort](java-quicksort)

This project is for STAT-441 Exprimental Design at
Montana State University - Bozeman during Spring 2019
as part of my PhD program.

## Running Experiment

1. Compile all sort algorithms for all the languages (treatments).
   - If you do not wish to run all of the tests then modify
     [`main.go`](./main.go) to indicate which tests are being run.
2. Use the following command to collect experimental data: `go run main.go`
3. Results are collected in `result.txt`

## Compiling Treatments

### Go

#### Go Binarytree

```sh
cd Go/Binarytree
go build
```

#### Go Mergesort

```sh
cd Go/Mergesort
go build
```

#### Go Quicksort

```sh
cd Go/Quicksort
go build
```

### Java

#### Java Binarytree

```sh
cd Java/Binarytree
javac Binarytree.java
```

#### Java Mergesort

```sh
cd Java/Mergesort
javac Mergesort.java
```

#### Java Quicksort

```sh
cd Java/Quicksort
javac Quicksort.java
```