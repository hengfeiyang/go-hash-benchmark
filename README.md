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
BenchmarkHash_fnv64-10                      188052156           6.301 ns/op        0 B/op          0 allocs/op
BenchmarkHash_fnv64_std-10                  141364040           8.540 ns/op        0 B/op          0 allocs/op
BenchmarkHash_crc64-10                      89487960            13.20 ns/op        0 B/op          0 allocs/op
BenchmarkHash_xxhash_cespare-10             165165523           7.204 ns/op        0 B/op          0 allocs/op
BenchmarkHash_xxhash_zeebo-10               34495850            34.04 ns/op        0 B/op          0 allocs/op
BenchmarkHash_xxhash_OneOfOne-10            99395344            11.52 ns/op        0 B/op          0 allocs/op
BenchmarkHash_cityhash_zentures-10          216331137           5.526 ns/op        0 B/op          0 allocs/op
BenchmarkHash_cityhash_tenfyzhong-10        220778608           5.441 ns/op        0 B/op          0 allocs/op
BenchmarkHash_cityhash_creachadair-10       238392873           5.036 ns/op        0 B/op          0 allocs/op
BenchmarkHash_murmur3_64-10                 52582378            21.53 ns/op       16 B/op          1 allocs/op
BenchmarkHash_farmhash64-10                 189131232           6.340 ns/op        0 B/op          0 allocs/op
BenchmarkHash_mmh3-10                       161781974           7.460 ns/op        0 B/op          0 allocs/op
BenchmarkHash_metrohash-10                  128223290           9.368 ns/op        0 B/op          0 allocs/op
BenchmarkHash_metro-10                      256281772           4.684 ns/op        0 B/op          0 allocs/op
BenchmarkHash_wyhash_orisano-10             45266292            25.35 ns/op       16 B/op          1 allocs/op
BenchmarkHash_wyhash_zhangyunhao116_v1-10   480504369           2.550 ns/op	       0 B/op          0 allocs/op
BenchmarkHash_wyhash_zhangyunhao116_v3-10   393447082           3.043 ns/op	       0 B/op          0 allocs/op
BenchmarkHash_highwayhash_minIO-10          29155610            40.53 ns/op	       0 B/op          0 allocs/op
BenchmarkHash_pengyhash2-10                 49536200            24.40 ns/op	       0 B/op          0 allocs/op
```

```
goos: linux
goarch: amd64
pkg: github.com/hengfeiyang/go-hash-benchmark
cpu: Intel(R) Xeon(R) Platinum 8375C CPU @ 2.90GHz
BenchmarkHash_fnv64-4                      	139189701	        8.619 ns/op	       0 B/op	       0 allocs/op
BenchmarkHash_fnv64_std-4                  	100000000	        11.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkHash_crc64-4                      	73295024	        16.26 ns/op	       0 B/op	       0 allocs/op
BenchmarkHash_xxhash_cespare-4             	61292402	        19.52 ns/op	       0 B/op	       0 allocs/op
BenchmarkHash_xxhash_zeebo-4               	43399722	        27.58 ns/op	       0 B/op	       0 allocs/op
BenchmarkHash_xxhash_OneOfOne-4            	54548454	        22.13 ns/op	       0 B/op	       0 allocs/op
BenchmarkHash_cityhash_zentures-4          	89041736	        13.23 ns/op	       0 B/op	       0 allocs/op
BenchmarkHash_cityhash_tenfyzhong-4        	87733195	        13.72 ns/op	       0 B/op	       0 allocs/op
BenchmarkHash_cityhash_creachadair-4       	87566337	        13.69 ns/op	       0 B/op	       0 allocs/op
BenchmarkHash_murmur3_64-4                 	34760779	        33.10 ns/op	      16 B/op	       1 allocs/op
BenchmarkHash_farmhash64-4                 	82134117	        13.99 ns/op	       0 B/op	       0 allocs/op
BenchmarkHash_mmh3-4                       	73028557	        15.82 ns/op	       0 B/op	       0 allocs/op
BenchmarkHash_metrohash-4                  	50959827	        23.56 ns/op	       0 B/op	       0 allocs/op
BenchmarkHash_metro-4                      	275654774	        4.353 ns/op	       0 B/op	       0 allocs/op
BenchmarkHash_wyhash_orisano-4             	29965324	        38.65 ns/op	      16 B/op	       1 allocs/op
BenchmarkHash_wyhash_zhangyunhao116_v1-4   	331706686	        3.559 ns/op	       0 B/op	       0 allocs/op
BenchmarkHash_wyhash_zhangyunhao116_v3-4   	266339889	        4.504 ns/op	       0 B/op	       0 allocs/op
BenchmarkHash_highwayhash_minIO-4          	26071016	        46.20 ns/op	       0 B/op	       0 allocs/op
BenchmarkHash_pengyhash2-4                 	28242212	        42.34 ns/op	       0 B/op	       0 allocs/op
```
