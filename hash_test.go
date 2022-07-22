package gohashbenchmark

import (
	"encoding/hex"
	"hash/crc64"
	"hash/fnv"
	"testing"

	"github.com/DataDog/mmh3"
	onexxhash "github.com/OneOfOne/xxhash"
	"github.com/cespare/xxhash/v2"
	crcityhash "github.com/creachadair/cityhash"
	"github.com/dgryski/go-metro"
	"github.com/hengfeiyang/go-hash-benchmark/farmhash64"
	"github.com/hengfeiyang/go-hash-benchmark/fnv64"
	"github.com/minio/highwayhash"
	"github.com/orisano/wyhash"
	"github.com/shivakar/metrohash"
	"github.com/skeeto/pengyhash"
	"github.com/spaolacci/murmur3"
	tcityhash "github.com/tenfyzhong/cityhash"
	"github.com/zeebo/xxh3"
	"github.com/zentures/cityhash"
	zwyhash "github.com/zhangyunhao116/wyhash"
)

var key = "1QvEL0YywgM"

var keyHash, _ = hex.DecodeString("000102030405060708090A0B0C0D0E0FF0E0D0C0B0A090807060504030201000") // use for minIO highwayHash

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

func BenchmarkHash_crc64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		table := crc64.MakeTable(crc64.ISO)
		crc64.Checksum([]byte(key), table)
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

func BenchmarkHash_murmur3_64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		murmur3.Sum64([]byte(key))
	}
}

func BenchmarkHash_farmhash64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		farmhash64.FarmHash64([]byte(key))
	}
}

func BenchmarkHash_mmh3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mmh3.Hash32([]byte(key))
	}
}

func BenchmarkHash_metrohash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h := metrohash.NewMetroHash64()
		h.Write([]byte(key))
		h.Sum64()
	}
}

func BenchmarkHash_metro(b *testing.B) {
	for i := 0; i < b.N; i++ {
		metro.Hash64Str(key, 0)
	}
}

func BenchmarkHash_wyhash_orisano(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wyhash.Sum64(0, []byte(key))
	}
}

func BenchmarkHash_wyhash_zhangyunhao116_v1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		zwyhash.Sum64String(key)
	}
}

func BenchmarkHash_wyhash_zhangyunhao116_v3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		zwyhash.Sum64StringV3(key)
	}
}

func BenchmarkHash_highwayhash_minIO(b *testing.B) {
	for i := 0; i < b.N; i++ {
		highwayhash.Sum64([]byte(key), keyHash)
	}
}

func BenchmarkHash_pengyhash2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pengyhash.Pengyhash([]byte(key), uint32(len(key)))
	}
}
