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
BenchmarkHash_meow_mmcloughlin-4           	 3246139	       365.9 ns/op	    1208 B/op	       6 allocs/op
BenchmarkHash_wyhash_orisano-4             	30461908	        38.48 ns/op	      16 B/op	       1 allocs/op
BenchmarkHash_wyhash_zhangyunhao116_v1-4   	332034903	         3.613 ns/op	       0 B/op	       0 allocs/op
BenchmarkHash_wyhash_zhangyunhao116_v3-4   	266124853	         4.511 ns/op	       0 B/op	       0 allocs/op
BenchmarkHash_highwayhash_minIO-4          	24491539	        49.18 ns/op	       0 B/op	       0 allocs/op
BenchmarkHash_pengyhash-4                  	18162232	        65.57 ns/op	      32 B/op	       1 allocs/op
BenchmarkHash_pengyhash2-4                 	27764001	        42.48 ns/op	       0 B/op	       0 allocs/op
```