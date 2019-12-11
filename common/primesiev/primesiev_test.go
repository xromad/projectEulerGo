package primesiev

import (
	"projectEuler/common/commonintutils"
	"testing"
)

func TestPrimeFactorInt64(t *testing.T) {
	testCases := []struct {
		number int64
		want   []int64
	}{
		{0, []int64{1}},
		{1, []int64{1}},
		{2, []int64{1, 2}},
		{3, []int64{1, 3}},
		{5, []int64{1, 5}},
		{7, []int64{1, 7}},
		{12, []int64{1, 2, 2, 3}},
		{15, []int64{1, 3, 5}},
		{13195, []int64{1, 5, 7, 13, 29}}, //the provided example
	}

	for _, c := range testCases {
		got := PrimeFactorInt64(c.number)
		if !commonintutils.Int64SliceEquals(got, c.want) {
			t.Errorf("PrimeFactorInt64(%v) == %v, want %v", c.number, got, c.want)
		}
	}
}

func TestIsPrime(t *testing.T) {
	testCases := []struct {
		primes []int64
		num    int64
		want   bool
	}{
		{[]int64{1, 2, 3}, 4, false},
		{[]int64{1, 2, 3}, 5, true},
	}

	for _, c := range testCases {
		got := isPrime(c.num, c.primes)
		if got != c.want {
			t.Errorf("isPrime(%v, %v) == %t, want %t", c.num, c.primes, got, c.want)
		}
	}
}

func TestFindPrimesSmallerThan(t *testing.T) {
	testCases := []struct {
		max  int64
		want []int64
	}{
		{3, []int64{1, 2, 3}},
		{6, []int64{1, 2, 3, 5}},
	}
	for _, c := range testCases {
		got := FindPrimesSmallerThan(c.max)
		if !commonintutils.Int64SliceEquals(c.want, got) {
			t.Errorf("FindPrimesSmallerThan(%v) == %v, want %v", c.max, got, c.want)
		}
	}
}

func TestGetNextPrime(t *testing.T) {
	testCases := []struct {
		primesIn []int64
		want     []int64
	}{
		{[]int64{}, []int64{1}},
		{[]int64{1}, []int64{1, 2}},
		{[]int64{1, 2, 3, 5, 7}, []int64{1, 2, 3, 5, 7, 11}},
	}
	for _, c := range testCases {
		got := GetNextPrime(&c.primesIn)
		if !commonintutils.Int64SliceEquals(c.want, got) {
			t.Errorf("GetnextPrime(%v) == %v, want %v", c.primesIn, got, c.want)
		}
	}
}

func TestFindPrimes(t *testing.T) {
	testCases := []struct {
		count int
		want  []int64
	}{
		{4, []int64{1, 2, 3, 5}},
		{5, []int64{1, 2, 3, 5, 7}},
	}
	for _, c := range testCases {
		got := FindPrimes(c.count)
		if !commonintutils.Int64SliceEquals(c.want, got) {
			t.Errorf("Findprimes(%d) == %v, want %v", c.count, got, c.want)
		}
	}
}

func TestFindPrimesFrom(t *testing.T) {
	testCases := []struct {
		count int
		start int
		want  []int64
	}{
		{0, 0, []int64{}},
		{0, 1, []int64{}},
		{1, 0, []int64{1}},
		{1, 1, []int64{1}},
		{1, 2, []int64{2}},
		{4, 1, []int64{1, 2, 3, 5}},
		{4, 2, []int64{2, 3, 5, 7}},
	}
	for _, c := range testCases {
		got := FindPrimesFrom(c.count, c.start)
		if !commonintutils.Int64SliceEquals(c.want, got) {
			t.Errorf("FindprimesFrom(%d, %d) == %v, want %v", c.count, c.start, got, c.want)
		}
	}
}
