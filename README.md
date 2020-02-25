# workers

example to create a worker pool using channels and goroutines to provide maximum throughput.

the code consists of 2 examples a_sequential where the work happens in sequence and b_concurrent where the work gets distributed
among several workers. The users can run the examples themselves and see the difference.


avg time 10 sequential runs := 17.25387443 Seconds
avg time 10 concurrent runs := 1.494497171 Seconds



Benchmarks for concurrent runs

### benchmark with 1 worker
```shell script
goos: linux
goarch: amd64
pkg: github.com/PrakharSrivastav/workers
BenchmarkConcurrent-8                  1        4349083504 ns/op
BenchmarkNonconcurrent-8               1        4276954083 ns/op
PASS
ok      github.com/PrakharSrivastav/workers     8.630s
```


### benchmark with 2 workers

```shell script
prakhar@tardis (master)✗ % go test -bench=.                                                                                                                                    ~/Workspace/examples/blog.examples/understanding-context
goos: linux
goarch: amd64
pkg: github.com/PrakharSrivastav/workers
BenchmarkConcurrent-8                  1        2341316053 ns/op
BenchmarkNonconcurrent-8               1        2076455734 ns/op
PASS
ok      github.com/PrakharSrivastav/workers     4.422s
```


### benchmark with 4 workers
```shell script
prakhar@tardis (master)✗ % go test -bench=.                                                                                                                                    ~/Workspace/examples/blog.examples/understanding-context
goos: linux
goarch: amd64
pkg: github.com/PrakharSrivastav/workers
BenchmarkConcurrent-8                  1        1318867706 ns/op
BenchmarkNonconcurrent-8               1        1076381168 ns/op
PASS
ok      github.com/PrakharSrivastav/workers     2.399s
```

### benchmark with 8 workers
```shell script
prakhar@tardis (master)✗ % go test -bench=.                                                                                                                                    ~/Workspace/examples/blog.examples/understanding-context
goos: linux
goarch: amd64
pkg: github.com/PrakharSrivastav/workers
BenchmarkConcurrent-8                  2         561313450 ns/op
BenchmarkNonconcurrent-8               2         544786952 ns/op
PASS
ok      github.com/PrakharSrivastav/workers     3.554s
```

### benchmark with 16 workers
```shell script
prakhar@tardis (master)✗ % go test -bench=.                                                                                                                                    ~/Workspace/examples/blog.examples/understanding-context
goos: linux
goarch: amd64
pkg: github.com/PrakharSrivastav/workers
BenchmarkConcurrent-8                  3         397863120 ns/op
BenchmarkNonconcurrent-8               3         367798299 ns/op
PASS
ok      github.com/PrakharSrivastav/workers     4.747s
```