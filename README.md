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
BenchmarkHash_fnv64-10                          188052156                6.301 ns/op           0 B/op          0 allocs/op
BenchmarkHash_fnv64_std-10                      141364040                8.540 ns/op           0 B/op          0 allocs/op
BenchmarkHash_fnv32_std-10                      141947005                8.405 ns/op           0 B/op          0 allocs/op
BenchmarkHash_crc64-10                          89487960                 13.20 ns/op           0 B/op          0 allocs/op
BenchmarkHash_xxhash_cespare-10                 165165523                7.204 ns/op           0 B/op          0 allocs/op
BenchmarkHash_xxhash_zeebo-10                   34495850                 34.04 ns/op           0 B/op          0 allocs/op
BenchmarkHash_xxhash_OneOfOne-10                99395344                 11.52 ns/op           0 B/op          0 allocs/op
BenchmarkHash_cityhash_zentures-10              216331137                5.526 ns/op           0 B/op          0 allocs/op
BenchmarkHash_cityhash_tenfyzhong-10            220778608                5.441 ns/op           0 B/op          0 allocs/op
BenchmarkHash_cityhash_creachadair-10           238392873                5.036 ns/op           0 B/op          0 allocs/op
BenchmarkHash_murmur3_64-10                     52582378                 21.53 ns/op          16 B/op          1 allocs/op
BenchmarkHash_farmhash64-10                     189131232                6.340 ns/op           0 B/op          0 allocs/op
BenchmarkHash_mmh3-10                           161781974                7.460 ns/op           0 B/op          0 allocs/op
BenchmarkHash_metrohash-10                      128223290                9.368 ns/op           0 B/op          0 allocs/op
BenchmarkHash_metro-10                          256281772                4.684 ns/op           0 B/op          0 allocs/op
BenchmarkHash_wyhash_orisano-10                 45266292                 25.35 ns/op          16 B/op          1 allocs/op
BenchmarkHash_wyhash_zhangyunhao116_v1-10       480504369                2.550 ns/op           0 B/op          0 allocs/op
BenchmarkHash_wyhash_zhangyunhao116_v3-10       393447082                3.043 ns/op           0 B/op          0 allocs/op
BenchmarkHash_highwayhash_minIO-10              29155610                 40.53 ns/op           0 B/op          0 allocs/op
BenchmarkHash_pengyhash2-10                     49536200                 24.40 ns/op           0 B/op          0 allocs/op
```

```
goos: linux
goarch: amd64
pkg: github.com/hengfeiyang/go-hash-benchmark
cpu: Intel(R) Xeon(R) Platinum 8375C CPU @ 2.90GHz
BenchmarkHash_fnv64-4                      	184197849	         6.505 ns/op	       0 B/op	       0 allocs/op
BenchmarkHash_fnv64_std-4                  	100000000	        10.26 ns/op	       0 B/op	       0 allocs/op
BenchmarkHash_fnv32_std-4                  	127840286	         9.397 ns/op	       0 B/op	       0 allocs/op
BenchmarkHash_crc32-4                      	30165579	        37.14 ns/op	      16 B/op	       1 allocs/op
BenchmarkHash_crc32_klauspost-4            	31531101	        36.87 ns/op	      16 B/op	       1 allocs/op
BenchmarkHash_crc64-4                      	64323390	        18.58 ns/op	       0 B/op	       0 allocs/op
BenchmarkHash_xxhash_cespare-4             	61138449	        19.59 ns/op	       0 B/op	       0 allocs/op
BenchmarkHash_xxhash_zeebo-4               	41720746	        27.46 ns/op	       0 B/op	       0 allocs/op
BenchmarkHash_xxhash_OneOfOne-4            	53066194	        22.35 ns/op	       0 B/op	       0 allocs/op
BenchmarkHash_cityhash_zentures-4          	88944982	        13.38 ns/op	       0 B/op	       0 allocs/op
BenchmarkHash_cityhash_tenfyzhong-4        	86932873	        13.47 ns/op	       0 B/op	       0 allocs/op
BenchmarkHash_cityhash_creachadair-4       	85694116	        13.99 ns/op	       0 B/op	       0 allocs/op
BenchmarkHash_murmur3_32-4                 	15559698	        75.55 ns/op	      96 B/op	       2 allocs/op
BenchmarkHash_murmur3_64-4                 	14272201	        82.67 ns/op	     112 B/op	       2 allocs/op
BenchmarkHash_farmhash64-4                 	84250557	        14.09 ns/op	       0 B/op	       0 allocs/op
BenchmarkHash_mmh3-4                       	74675529	        16.02 ns/op	       0 B/op	       0 allocs/op
BenchmarkHash_metrohash-4                  	50637700	        23.60 ns/op	       0 B/op	       0 allocs/op
BenchmarkHash_metro-4                      	275270512	         4.358 ns/op	       0 B/op	       0 allocs/op
BenchmarkHash_wyhash_orisano-4             	30461908	        38.48 ns/op	      16 B/op	       1 allocs/op
BenchmarkHash_wyhash_zhangyunhao116_v1-4   	332034903	         3.613 ns/op	       0 B/op	       0 allocs/op
BenchmarkHash_wyhash_zhangyunhao116_v3-4   	266124853	         4.511 ns/op	       0 B/op	       0 allocs/op
BenchmarkHash_highwayhash_minIO-4          	24491539	        49.18 ns/op	       0 B/op	       0 allocs/op
BenchmarkHash_pengyhash-4                  	18162232	        65.57 ns/op	      32 B/op	       1 allocs/op
BenchmarkHash_pengyhash2-4                 	27764001	        42.48 ns/op	       0 B/op	       0 allocs/op
```