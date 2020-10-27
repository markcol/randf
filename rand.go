// rand.go -- Random number generation
//
// The goal of this algorithm is to generate uniformly-distributed pseudo-random
// floating-point values in the range [0, 1], assuming that we are given an algorithm
// for generating pseudo-random bits. These bits should have probability 50% of
// being in each of two states, and be statistically independent.
//
// The approach is to choose floating-point values in the range such that the
// probability that a given value is chosen is proportional to the distance between
// it and its two neighbors. The algorithm works in two steps, first choosing the
// exponent range of the value, then choosing the mantissa.
//
// From the paper http://allendowney.com/research/rand/downey07randfloat.pdf.

package randf

import (
	"math"
	"math/rand"
)

var (
	lowExp  = (math.Float32bits(0.0) >> 23) & 0xFF // exponent field of 0.0
	highExp = (math.Float32bits(1.0) >> 23) & 0xFF // exponent field of 1.0
)

// A Rand is a sourcce of random numbers.
type Rand struct {
	bits int32       // holds the number of bits remaining
	x    int32       // holds the value of bits remaining
	src  rand.Source // our random number source
	rand *rand.Rand  // our random number
}

// New returns a new Rand that uses random values to generate other random values.
func New() *Rand {
	src := rand.NewSource(1)
	return &Rand{
		src:  src,
		rand: rand.New(src),
	}
}

// Seed uses the provided seed value to initialize the default Source to a
// deterministic state. If Seed is not called, the generator behaves as if
// seeded by Seed(1). Seed values that have the same remainder when divided
// by 2³¹-1 generate the same pseudo-random sequence. Seed is safe for
// concurrent use.
func (r *Rand) Seed(v int64) {
	r.rand.Seed(v)
}

// Float32 returns a random 32-bit floating-point number in the range
// [0.0,1.0], including 0.0, subnormals, and 1.0.
func (r *Rand) Float32() float32 {
	var exp uint32

	// Choose random bits and decrement exp until a 1 appears.
	for exp = highExp - 1; exp > lowExp; exp-- {
		if b := r.getBit(); b != 0 {
			break
		}
	}

	// Choose a random 23-bit mantissa
	mant := uint32(r.rand.Int31() & 0x7FFFFF)

	// If the mantissa is zero, half the time we should move to the next
	// exponent range
	if mant == 0 && r.getBit() != 0 {
		exp++
	}

	// Combine the exponent and the mantissa
	return math.Float32frombits((exp << 23) | mant)
}

// getBit returns a random bit. For efficiency, bits are generated 31 at a
// time using the function rand.Int31().
func (r *Rand) getBit() int32 {
	if r.bits == 0 {
		r.x = r.rand.Int31()
		r.bits = 31
	}
	bit := r.x & 1
	r.x >>= 1
	r.bits--
	return bit
}
