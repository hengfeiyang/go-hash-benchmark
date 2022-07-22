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
BenchmarkHash_fnv64-10                      188052156           6.301 ns/op           0 B/op           0 allocs/op
BenchmarkHash_fnv64_std-10                  141364040           8.540 ns/op           0 B/op           0 allocs/op
BenchmarkHash_crc64-10                      89487960            13.20 ns/op           0 B/op           0 allocs/op
BenchmarkHash_xxhash_cespare-10             165165523           7.204 ns/op           0 B/op           0 allocs/op
BenchmarkHash_xxhash_zeebo-10               34495850            34.04 ns/op           0 B/op           0 allocs/op
BenchmarkHash_xxhash_OneOfOne-10            99395344            11.52 ns/op           0 B/op           0 allocs/op
BenchmarkHash_cityhash_zentures-10          216331137           5.526 ns/op           0 B/op           0 allocs/op
BenchmarkHash_cityhash_tenfyzhong-10        220778608           5.441 ns/op           0 B/op           0 allocs/op
BenchmarkHash_cityhash_creachadair-10       238392873           5.036 ns/op           0 B/op           0 allocs/op
BenchmarkHash_murmur3_64-10                 52582378            21.53 ns/op          16 B/op           1 allocs/op
BenchmarkHash_farmhash64-10                 189131232           6.340 ns/op           0 B/op           0 allocs/op
BenchmarkHash_mmh3-10                       161781974           7.460 ns/op           0 B/op           0 allocs/op
BenchmarkHash_metrohash-10                  128223290           9.368 ns/op           0 B/op           0 allocs/op
BenchmarkHash_metro-10                      256281772           4.684 ns/op           0 B/op           0 allocs/op
BenchmarkHash_wyhash_orisano-10             45266292            25.35 ns/op          16 B/op           1 allocs/op
BenchmarkHash_wyhash_zhangyunhao116_v1-10   480504369           2.550 ns/op           0 B/op           0 allocs/op
BenchmarkHash_wyhash_zhangyunhao116_v3-10   393447082           3.043 ns/op           0 B/op           0 allocs/op
BenchmarkHash_highwayhash_minIO-10          29155610            40.53 ns/op           0 B/op           0 allocs/op
BenchmarkHash_pengyhash2-10                 49536200            24.40 ns/op           0 B/op           0 allocs/op
```

```
goos: linux
goarch: amd64
pkg: github.com/hengfeiyang/go-hash-benchmark
cpu: Intel(R) Xeon(R) Platinum 8375C CPU @ 2.90GHz
BenchmarkHash_fnv64-4                       185279481           6.484 ns/op           0 B/op           0 allocs/op
BenchmarkHash_fnv64_std-4                   100000000           10.31 ns/op           0 B/op           0 allocs/op
BenchmarkHash_crc64-4                       72141110            16.62 ns/op           0 B/op           0 allocs/op
BenchmarkHash_xxhash_cespare-4              60971256            19.59 ns/op           0 B/op           0 allocs/op
BenchmarkHash_xxhash_zeebo-4                43381956            27.56 ns/op           0 B/op           0 allocs/op
BenchmarkHash_xxhash_OneOfOne-4             53699167            22.34 ns/op           0 B/op           0 allocs/op
BenchmarkHash_cityhash_zentures-4           89680816            13.35 ns/op           0 B/op           0 allocs/op
BenchmarkHash_cityhash_tenfyzhong-4         87935768            13.63 ns/op           0 B/op           0 allocs/op
BenchmarkHash_cityhash_creachadair-4        87616258            13.62 ns/op           0 B/op           0 allocs/op
BenchmarkHash_murmur3_64-4                  34575916            33.62 ns/op          16 B/op           1 allocs/op
BenchmarkHash_farmhash64-4                  85884194            13.62 ns/op           0 B/op           0 allocs/op
BenchmarkHash_mmh3-4                        75253887            15.92 ns/op           0 B/op           0 allocs/op
BenchmarkHash_metrohash-4                   50896293            23.58 ns/op           0 B/op           0 allocs/op
BenchmarkHash_metro-4                       261343096           4.587 ns/op           0 B/op           0 allocs/op
BenchmarkHash_wyhash_orisano-4              30477706            38.30 ns/op          16 B/op           1 allocs/op
BenchmarkHash_wyhash_zhangyunhao116_v1-4    347594124           3.453 ns/op           0 B/op           0 allocs/op
BenchmarkHash_wyhash_zhangyunhao116_v3-4    266687262           4.482 ns/op           0 B/op           0 allocs/op
BenchmarkHash_highwayhash_minIO-4           26399698            45.81 ns/op           0 B/op           0 allocs/op
BenchmarkHash_pengyhash2-4                  28200452            42.38 ns/op           0 B/op           0 allocs/op
```
