# go-hash-benchmark
Benchmark for Hash function of Go

- fnv64
- xxhash
- cityhash

## result

```
goos: darwin
goarch: arm64
pkg: github.com/hengfeiyang/go-hash-benchmark
BenchmarkHash_fnv64-10                          189053720                6.226 ns/op           0 B/op          0 allocs/op
BenchmarkHash_fnv64_std-10                      142792552                8.399 ns/op           0 B/op          0 allocs/op
BenchmarkHash_fnv32_std-10                      142806360                8.402 ns/op           0 B/op          0 allocs/op
BenchmarkHash_xxhash_cespare-10                 167370736                7.162 ns/op           0 B/op          0 allocs/op
BenchmarkHash_xxhash_zeebo-10                   33958860                34.22 ns/op            0 B/op          0 allocs/op
BenchmarkHash_xxhash_OneOfOne-10                100000000               11.48 ns/op            0 B/op          0 allocs/op
BenchmarkHash_cityhash_zentures-10              217240786                5.517 ns/op           0 B/op          0 allocs/op
BenchmarkHash_cityhash_tenfyzhong-10            220452866                5.438 ns/op           0 B/op          0 allocs/op
BenchmarkHash_cityhash_creachadair-10           239254263                5.012 ns/op           0 B/op          0 allocs/op
```
