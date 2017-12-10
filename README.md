# gonum_compare

compare blas go implementation and OpenBlas go binding

```
BenchmarkDotFloat32ThousandGo   10000000           153 ns/op
BenchmarkDotFloat32ThousandCgo  10000000           160 ns/op
BenchmarkDotFloat32MillionGo        2000        556859 ns/op
BenchmarkDotFloat32MillionCgo       3000        537727 ns/op
BenchmarkDotFloat64ThousandGo    5000000           275 ns/op
BenchmarkDotFloat64ThousandCgo  10000000           216 ns/op
BenchmarkDotFloat64MillionGo        2000       1204199 ns/op
BenchmarkDotFloat64MillionCgo       2000       1163406 ns/op
PASS
ok      _/go/gonum_compare  16.382s
```
