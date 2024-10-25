# go-benchstat-demo
A demo for using `benchstat` to compare benchmark results, 
allowing us to continuously monitor performance throughout version changes.

## Usage
I've already tagged two versions for this demo, `v1.0.0` and `v2.0.0`. You can compare the benchmark results between these two versions by running the following command:
```shell
git pull git@github.com:iFurySt/go-benchstat-demo.git
cd go-benchstat-demo
git checkout v1.0.0
go test -bench=. -benchmem -count=10 ./... > old-`git rev-parse HEAD`.txt
git checkout v2.0.0
go test -bench=. -benchmem -count=10 ./... > new-`git rev-parse HEAD`.txt
benchstat old-*.txt new-*.txt > bench-results.txt
cat bench-results.txt
```

## References
- [benchstat](https://pkg.go.dev/golang.org/x/perf/cmd/benchstat)
