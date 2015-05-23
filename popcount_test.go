package popcount

import (
	"math/rand"

	"testing"
)

type TestEntry32 struct {
	n uint32
	c uint8
}

var testTable32 = []TestEntry32{
	{0x0, 0},
	{0x1, 1},
	{0x2, 1},
	{0x4, 1},
	{0x80000000, 1},
	{0x80000001, 2},
	{0xFFFFFFFF, 32},
	{0xA0A0A0A0, 8},
	{0xAAAAAAAA, 16},
	{0x50505050, 8},
	{0x55555555, 16},
}

type TestEntry64 struct {
	n uint64
	c uint8
}

var testTable64 = []TestEntry64{
	{0x0, 0},
	{0x1, 1},
	{0x2, 1},
	{0x4, 1},
	{0x8000000080000000, 2},
	{0x8000000180000001, 4},
	{0xFFFFFFFFFFFFFFFF, 64},
	{0xA0A0A0A0A0A0A0A0, 16},
	{0xAAAAAAAAAAAAAAAA, 32},
	{0x5050505050505050, 16},
	{0x5555555555555555, 32},
}

func testImpl32(t *testing.T, impl func(n uint32) uint8) {
	for _, test := range testTable32 {
		actual := impl(test.n)
		if actual != test.c {
			t.Errorf("test: %d expected: %d got: %d", test.n, test.c, actual)
		}
	}
}

func benchmarkImpl32(b *testing.B, impl func(n uint32) uint8) {
	var randomTable = [64]TestEntry32{}
	for i := 0; i < len(randomTable); i += 1 {
		n := rand.Uint32()
		randomTable[i] = TestEntry32{n, Slow32(n)}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		impl(randomTable[i%len(randomTable)].n)
	}
}

func testImpl64(t *testing.T, impl func(n uint64) uint8) {
	for _, test := range testTable64 {
		actual := impl(test.n)
		if actual != test.c {
			t.Errorf("test: %d expected: %d got: %d", test.n, test.c, actual)
		}
	}
}

func benchmarkImpl64(b *testing.B, impl func(n uint64) uint8) {
	var randomTable = [64]TestEntry64{}
	for i := 0; i < len(randomTable); i += 1 {
		n := uint64(rand.Uint32())<<32 | uint64(rand.Uint32())
		randomTable[i] = TestEntry64{n, Slow64(n)}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		impl(randomTable[i%len(randomTable)].n)
	}
}

func TestHamming32(t *testing.T) {
	testImpl32(t, Hamming32)
}

func BenchmarkHamming32(b *testing.B) {
	benchmarkImpl32(b, Hamming32)
}

func TestSlow32(t *testing.T) {
	testImpl32(t, Slow32)
}

func BenchmarkSlow32(b *testing.B) {
	benchmarkImpl32(b, Slow32)
}

func TestFast32(t *testing.T) {
	testImpl32(t, Fast32)
}

func BenchmarkFast32(b *testing.B) {
	benchmarkImpl32(b, Fast32)
}

func TestPopCnt32(t *testing.T) {
	testImpl32(t, PopCnt32)
}

func BenchmarkPopCnt32(b *testing.B) {
	benchmarkImpl32(b, PopCnt32)
}

func TestPopCnt64(t *testing.T) {
	testImpl64(t, PopCnt64)
}

func BenchmarkPopCnt64(b *testing.B) {
	benchmarkImpl64(b, PopCnt64)
}

func TestFast64(t *testing.T) {
	testImpl64(t, Fast64)
}

func BenchmarkFast64(b *testing.B) {
	benchmarkImpl64(b, Fast64)
}

func TestByteTable32(t *testing.T) {
	testImpl32(t, ByteTable32)
}

func BenchmarkByteTable32(b *testing.B) {
	benchmarkImpl32(b, ByteTable32)
}

func TestByteTable64(t *testing.T) {
	testImpl64(t, ByteTable64)
}

func BenchmarkByteTable64(b *testing.B) {
	benchmarkImpl64(b, ByteTable64)
}

func TestHamming64(t *testing.T) {
	testImpl64(t, Hamming64)
}

func BenchmarkHamming64(b *testing.B) {
	benchmarkImpl64(b, Hamming64)
}

func TestSlow64(t *testing.T) {
	testImpl64(t, Slow64)
}

func BenchmarkSlow64(b *testing.B) {
	benchmarkImpl64(b, Slow64)
}
