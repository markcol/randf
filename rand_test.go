package randf

import (
	"testing"
)

const maxReps = 10_000_000

// Ensure that using a default Rand is the same as using one with a seed that
// is explicitly set to 0.
func TestDefaultSeed(t *testing.T) {
	r1 := New()
	r2 := New()
	r2.Seed(1)
	for i := 0; i < maxReps; i++ {
		got := r1.Float32()
		expect := r2.Float32()
		if got != expect {
			t.Errorf("iteration %d: expected: %v, got %v", i, expect, got)
		}
	}
}

var result float32

func benchmarkFloat32(i int, b *testing.B) {
	var res float32
	r := New()
	for n := b.N; n > 0; n-- {
		// Always record the result of Float32() to prevent the compiler
		// eliminating the function call.
		res = r.Float32()
	}
	// Always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = res
}

func BenchmarkFloat1(b *testing.B)  { benchmarkFloat32(1, b) }
func BenchmarkFloat2(b *testing.B)  { benchmarkFloat32(2, b) }
func BenchmarkFloat3(b *testing.B)  { benchmarkFloat32(3, b) }
func BenchmarkFloat10(b *testing.B) { benchmarkFloat32(10, b) }
func BenchmarkFloat20(b *testing.B) { benchmarkFloat32(20, b) }
func BenchmarkFloat40(b *testing.B) { benchmarkFloat32(40, b) }
