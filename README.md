# Code Experiment

:seedling: **This experiment is not ready yet! Come back later.**

-------

This project is for STAT-441 Exprimental Design at
Montana State University - Bozeman during Spring 2019
as part of my PhD program.

## Running Experiment

1. Use the following command to build all of the sort
   algorithms for all the languages (treatments):
  - `go run build\build.go`
  - If you do not wish to build all of the treatments, then first modify
    [`build.go`](./build/build.go) to indicate which treatments are being built.
2. Use the following command to collect experimental data:
  - `go run main.go`
  - If you do not wish to run all of the tests, then first modify
    [`main.go`](./main.go) to indicate which tests are being run.
3. Results are collected in `result.txt`
