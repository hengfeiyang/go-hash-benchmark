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
BenchmarkHash_fnv64-10                          187190166                6.308 ns/op           0 B/op          0 allocs/op
BenchmarkHash_fnv64_std-10                      142402866                8.443 ns/op           0 B/op          0 allocs/op
BenchmarkHash_fnv32_std-10                      142227180                8.476 ns/op           0 B/op          0 allocs/op
BenchmarkHash_crc32-10                          60172746                18.57 ns/op           16 B/op          1 allocs/op
BenchmarkHash_crc32_klauspost-10                45879231                25.28 ns/op           16 B/op          1 allocs/op
BenchmarkHash_crc64-10                          90145645                12.80 ns/op            0 B/op          0 allocs/op
BenchmarkHash_xxhash_cespare-10                 167008308                7.154 ns/op           0 B/op          0 allocs/op
BenchmarkHash_xxhash_zeebo-10                   34723226                34.00 ns/op            0 B/op          0 allocs/op
BenchmarkHash_xxhash_OneOfOne-10                100000000               11.54 ns/op            0 B/op          0 allocs/op
BenchmarkHash_cityhash_zentures-10              216742515                5.522 ns/op           0 B/op          0 allocs/op
BenchmarkHash_cityhash_tenfyzhong-10            218083118                5.434 ns/op           0 B/op          0 allocs/op
BenchmarkHash_cityhash_creachadair-10           239661058                5.001 ns/op           0 B/op          0 allocs/op
BenchmarkHash_murmur3_32-10                     21377750                54.23 ns/op           96 B/op          2 allocs/op
BenchmarkHash_murmur3_64-10                     22351658                51.69 ns/op          112 B/op          2 allocs/op
BenchmarkHash_farmhash64-10                     189320271                6.339 ns/op           0 B/op          0 allocs/op
BenchmarkHash_mmh3-10                           161496896                7.415 ns/op           0 B/op          0 allocs/op
BenchmarkHash_metrohash-10                      127875795                9.368 ns/op           0 B/op          0 allocs/op
BenchmarkHash_metro-10                          256376679                4.682 ns/op           0 B/op          0 allocs/op
BenchmarkHash_meow_mmcloughlin-10                3102812               383.9 ns/op          1208 B/op          6 allocs/op
BenchmarkHash_wyhash_orisano-10                 46054356                25.44 ns/op           16 B/op          1 allocs/op
BenchmarkHash_wyhash_zhangyunhao116_v1-10       480207448                2.490 ns/op           0 B/op          0 allocs/op
BenchmarkHash_wyhash_zhangyunhao116_v3-10       390571018                3.069 ns/op           0 B/op          0 allocs/op
BenchmarkHash_highwayhash_minIO-10              29391714                40.32 ns/op            0 B/op          0 allocs/op
BenchmarkHash_pengyhash-10                      26482758                43.79 ns/op           32 B/op          1 allocs/op
BenchmarkHash_pengyhash2-10                     48789164                23.85 ns/op            0 B/op          0 allocs/op
```
