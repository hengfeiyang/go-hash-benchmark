package gohashbenchmark

import (
	"hash/fnv"
	"testing"

	onexxhash "github.com/OneOfOne/xxhash"
	"github.com/cespare/xxhash/v2"
	crcityhash "github.com/creachadair/cityhash"
	"github.com/hengfeiyang/go-hash-benchmark/fnv64"
	tcityhash "github.com/tenfyzhong/cityhash"
	"github.com/zeebo/xxh3"
	"github.com/zentures/cityhash"
)

var key = "1QvEL0YywgM"

func BenchmarkHash_fnv64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fnv64.HASH.Sum64(key)
	}
}

func BenchmarkHash_fnv64_std(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h := fnv.New64a()
		h.Write([]byte(key))
		h.Sum64()
	}
}

func BenchmarkHash_fnv32_std(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h := fnv.New32a()
		h.Write([]byte(key))
		h.Sum32()
	}
}

func BenchmarkHash_xxhash_cespare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xxhash.Sum64([]byte(key))
	}
}

func BenchmarkHash_xxhash_zeebo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h := xxh3.New()
		h.WriteString(key)
		h.Sum64()
	}
}

func BenchmarkHash_xxhash_OneOfOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h := onexxhash.New64()
		h.WriteString(key)
		h.Sum64()
	}
}

func BenchmarkHash_cityhash_zentures(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cityhash.CityHash64([]byte(key), uint32(len(key)))
	}
}

func BenchmarkHash_cityhash_tenfyzhong(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tcityhash.CityHash64([]byte(key))
	}
}

func BenchmarkHash_cityhash_creachadair(b *testing.B) {
	for i := 0; i < b.N; i++ {
		crcityhash.Hash64([]byte(key))
	}
}
