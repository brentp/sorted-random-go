// generate sorted random numbers
package sortedrandom

import (
	"math"
	"math/rand"
)

// Create a slice with increasing (sorted) random floats between 0 and 1.
func Float64(N int) []float64 {
	x := make([]float64, N)
	curmax := float64(1)
	for i := N; i > 0; i-- {
		curmax *= math.Pow(rand.Float64(), float64(1)/float64(i))
		x[i-1] = curmax
	}
	return x
}

// Generate increasing (sorted) random values.
func GenFloat64(N int) chan float64 {
	ch := make(chan float64, 4)
	go func() {
		curmax := float64(1)
		for i := N; i > 0; i-- {
			curmax *= math.Pow(rand.Float64(), float64(1)/float64(i))
			ch <- 1.0 - curmax
		}
		close(ch)
	}()
	return ch
}

// Create a slice with increasing (sorted) random floats between imin and imax.
func Int64(N int, imin int, imax int) []int64 {
	span := float64(imax - imin + 1)
	ints := make([]int64, N)

	for i, f := range Float64(N) {
		ints[i] = int64(imin) + int64(f*span)
	}
	return ints
}

// Generate increasing (sorted) random int64's between imin and imax.
func GenInt64(N int, imin int, imax int) chan int64 {
	ch := make(chan int64, 4)
	span := float64(imax - imin + 1)
	go func() {
		for f := range GenFloat64(N) {
			ch <- int64(imin) + int64(f*span)
		}
		close(ch)
	}()
	return ch
}
