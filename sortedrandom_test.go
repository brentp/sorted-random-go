package sortedrandom

import "testing"

func TestFloat64(t *testing.T) {
	N := 10
	flts := Float64(N)
	if len(flts) != N {
		t.Error("Float64 didn't return the correct length")
	}

	for i := 0; i < N-2; i++ {
		if flts[i] > flts[i+1] {
			t.Error("Should be increasing", flts)
		}
	}
}

func TestInt64(t *testing.T) {
	N := 10
	xmin := 12345
	xmax := 99999
	ints := Int64(N, xmin, xmax)
	if len(ints) != N {
		t.Error("Int64 didn't return the correct length")
	}

	for i := 0; i < N-2; i++ {
		if ints[i] > ints[i+1] {
			t.Error("Int64 Should be increasing", ints)
		}
	}
	if ints[0] < int64(xmin) {
		t.Error("First entry in int64 range should be >= xmin.")
	}
	if ints[N-1] > int64(xmax) {
		t.Error("Last entry in int64 range should be <= xmax.")
	}
}

func TestGenFloat64(t *testing.T) {
	fmin := float64(0)
	N := 10

	j := 0
	for f := range GenFloat64(N) {
		if fmin > f {
			t.Error("GenFloat64 should be increasing.")
		}
		fmin = f
		j += 1
	}
	if j != N {
		t.Error("GenFloat64 didn't generate the correct number of floats.")
	}

}

func TestGenInt64(t *testing.T) {
	imin := int64(0)
	N := 10

	j := 0
	for i := range GenInt64(N, 1234, 5555) {

		if imin > i {
			t.Error("IntFloat64 should be increasing.")
		}
		if i < 1234 {
			t.Error("Values from GenInt64 should be > imin.")
		}
		if i > 5555 {
			t.Error("Values from GenInt64 should be < imax.")
		}

		imin = i
		j += 1
	}
	if j != N {
		t.Error("GenFloat64 didn't generate the correct number of floats.")
	}

}
