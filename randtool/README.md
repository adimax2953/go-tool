# randtool#

randtool is a thread safe based on rand. It was originally implemented in golang by gosl. Due to lightweight requirements, so only copy the some rnd function from gosl.


## Install

```console
go get -u -v github.com/adimax2953/go-tool/randtool
```

## Usage

Let's start with a trivial example:

```go
package main

import (
	"github.com/adimax2953/go-tool/randtool"
	"fmt"
)

func main() {
    r := randtool.New(rand.NewSource(4872))
	r.Seed(4872)

	println("When seed is 4872. Int63r(0,10) return", r.Int63r(0, 10), ".")
}
```
## Benckmark

```
go test -benchmem -bench .
goos: windows
goarch: amd64
pkg: github.com/adimax2953/go-tool/randtool
cpu: 12th Gen Intel(R) Core(TM) i7-1260P
Benchmark_rand_Int63Threadsafe-16               99378988                12.57 ns/op            0 B/op          0 allocs/op
Benchmark_mt_Int63Threadsafe-16                 85725919                12.91 ns/op            0 B/op          0 allocs/op
Benchmark_rand_Int63ThreadsafeParallel-16       19362708                61.10 ns/op            0 B/op          0 allocs/op
Benchmark_rand_Int63Unthreadsafe-16             85503579                12.70 ns/op            0 B/op          0 allocs/op
Benchmark_mt_Int63ThreadsafeParallel-16         16552088                60.95 ns/op            0 B/op          0 allocs/op
Benchmark_rand_New-16                             136260              8189 ns/op               0 B/op          0 allocs/op
Benchmark_mt_New-16                             55566102                18.68 ns/op            0 B/op          0 allocs/op
Benchmark_rand_Intn1000-16                      80036816                14.03 ns/op            0 B/op          0 allocs/op
Benchmark_mt_Intn1000-16                        74885331                13.67 ns/op            0 B/op          0 allocs/op
Benchmark_rand_Int63n1000-16                    85495660                12.51 ns/op            0 B/op          0 allocs/op
Benchmark_mt_Int63n1000-16                      74875051                13.56 ns/op            0 B/op          0 allocs/op
Benchmark_rand_Int31n1000-16                    85672672                12.32 ns/op            0 B/op          0 allocs/op
Benchmark_mt_Int31n1000-16                      74610160                13.70 ns/op            0 B/op          0 allocs/op
Benchmark_rand_Float32-16                       105455782               12.10 ns/op            0 B/op          0 allocs/op
Benchmark_mt_Float32-16                         79111315                12.67 ns/op            0 B/op          0 allocs/op
Benchmark_rand_Float64-16                       85725919                11.94 ns/op            0 B/op          0 allocs/op
Benchmark_mt_Float64-16                         93917164                12.21 ns/op            0 B/op          0 allocs/op
Benchmark_rand_Perm3-16                         35260750                32.91 ns/op           24 B/op          1 allocs/op
Benchmark_mt_Perm3-16                           27004096                38.24 ns/op           24 B/op          1 allocs/op
Benchmark_rand_Perm30-16                         6019870               198.9 ns/op           240 B/op          1 allocs/op
Benchmark_mt_Perm30-16                           4984977               230.4 ns/op           240 B/op          1 allocs/op
Benchmark_rand_Perm30ViaShuffle-16              10000416               119.4 ns/op             0 B/op          0 allocs/op
Benchmark_mt_Perm30ViaShuffle-16                 6696027               165.1 ns/op             0 B/op          0 allocs/op
Benchmark_rand_ShuffleOverhead-16                6019665               197.8 ns/op             0 B/op          0 allocs/op
Benchmark_mt_ShuffleOverhead-16                  4278529               274.1 ns/op             0 B/op          0 allocs/op
Benchmark_rand_Read3-16                         74887200                13.54 ns/op            0 B/op          0 allocs/op
Benchmark_mt_Read3-16                           74952217                14.10 ns/op            0 B/op          0 allocs/op
Benchmark_rand_Read64-16                        24815178                42.52 ns/op            0 B/op          0 allocs/op
Benchmark_mt_Read64-16                          18323883                61.53 ns/op            0 B/op          0 allocs/op
Benchmark_rand_Read1000-16                       2666455               474.4 ns/op             0 B/op          0 allocs/op
Benchmark_mt_Read1000-16                         1274896               874.1 ns/op             0 B/op          0 allocs/op
Benchmark_rand_Int63r1000-16                    85711836                12.49 ns/op            0 B/op          0 allocs/op
Benchmark_mt_Int63r1000-16                      100971969               13.01 ns/op            0 B/op          0 allocs/op
Benchmark_rand_Int63s30-16                       3398407               344.7 ns/op             0 B/op          0 allocs/op
Benchmark_mt_Int63s30-16                         2921733               409.5 ns/op             0 B/op          0 allocs/op
ok      github.com/adimax2953/go-tool/randtool  112.697s



```




----------

## TODO

1. [X] Add TG Us test method.
2. [X] Improve or remove useless code.
3. [ ] Check code formatting.
