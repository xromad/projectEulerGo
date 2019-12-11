package commonintutils

import (
	"projectEuler/common/commonstringutils"
	"testing"
)

func TestProductIntSlice(t *testing.T) {
	testCases := []struct {
		intSlice []int
		want     int
	}{
		{[]int{}, 0},
		{[]int{0}, 0},
		{[]int{1}, 1},
		{[]int{1, 2}, 2},
		{[]int{1, 2, 3}, 6},
	}

	for _, c := range testCases {
		got := ProductIntSlice(c.intSlice)
		if got != c.want {
			t.Errorf("ProductIntSlice(%v) == %v, want %v", c.intSlice, got, c.want)
		}
	}
}

func TestI64ToA(t *testing.T) {
	testCases := []struct {
		num  int64
		want string
	}{
		{0, "0"},
		{1001, "1001"},
	}

	for _, c := range testCases {
		got := I64toA(c.num)
		if got != c.want {
			t.Errorf("I64toA(%v) == %s, want %s", c.num, got, c.want)
		}
	}
}

func TestIntSliceToStringSlice(t *testing.T) {
	testCases := []struct {
		intSlice []int
		want     []string
	}{
		{[]int{}, []string{}},
		{[]int{3, 5, 6, 9}, []string{"3", "5", "6", "9"}},
	}

	for _, c := range testCases {
		got := IntSliceToStringSlice(c.intSlice)
		if !commonstringutils.StringSliceCompare(got, c.want) {
			t.Errorf("IntSliceToStringSlice(%d) == %s, want %s", c.intSlice, got, c.want)
		}
	}
}

func TestSumIntSlice(t *testing.T) {
	testCases := []struct {
		intSlice []int
		want     int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 2}, 3},
	}

	for _, c := range testCases {
		got := SumIntSlice(c.intSlice)
		if got != c.want {
			t.Errorf("SumIntSlice(%v) == %v, want %v", c.intSlice, got, c.want)
		}
	}
}

func TestSumInt64Slice(t *testing.T) {
	testCases := []struct {
		intSlice []int64
		want     int64
	}{
		{[]int64{}, 0},
		{[]int64{1}, 1},
		{[]int64{1, 2}, 3},
	}

	for _, c := range testCases {
		got := SumInt64Slice(c.intSlice)
		if got != c.want {
			t.Errorf("SumInt64Slice(%v) == %v, want %v", c.intSlice, got, c.want)
		}
	}
}

func TestIntSliceContains(t *testing.T) {
	testCases := []struct {
		sliceA []int
		s      int
		want   bool
	}{
		{[]int{}, 1, false},
		{[]int{1}, 1, true},
		{[]int{1}, 2, false},
		{[]int{1, 2}, 1, true},
		{[]int{1, 2}, 2, true},
	}

	for _, c := range testCases {
		got := IntSliceContains(c.sliceA, c.s)
		if got != c.want {
			t.Errorf("IntSliceContains(%v, %v) == %v, want %v", c.sliceA, c.s, got, c.want)
		}
	}
}

func TestInt64SliceContains(t *testing.T) {
	testCases := []struct {
		sliceA []int64
		s      int64
		want   bool
	}{
		{[]int64{}, 1, false},
		{[]int64{1}, 1, true},
		{[]int64{1}, 2, false},
		{[]int64{1, 2}, 1, true},
		{[]int64{1, 2}, 2, true},
	}

	for _, c := range testCases {
		got := Int64SliceContains(c.sliceA, c.s)
		if got != c.want {
			t.Errorf("Int64SliceContains(%v, %v) == %v, want %v", c.sliceA, c.s, got, c.want)
		}
	}
}

func TestIntSliceEquals(t *testing.T) {
	testCases := []struct {
		sliceA []int
		sliceB []int
		want   bool
	}{
		{[]int{}, []int{}, true},
		{[]int{1}, []int{}, false},
		{[]int{}, []int{1}, false},
		{[]int{1}, []int{1}, true},
		{[]int{1, 2}, []int{1}, false},
		{[]int{1, 2}, []int{2}, false},
		{[]int{1}, []int{1, 2}, false},
		{[]int{2}, []int{1, 2}, false},
		{[]int{1, 2}, []int{1, 2}, true},
		{[]int{1, 2}, []int{1, 3}, false},
	}

	for _, c := range testCases {
		got := IntSliceEquals(c.sliceA, c.sliceB)
		if got != c.want {
			t.Errorf("IntSliceEquals(%v, %v) == %v, want %v", c.sliceA, c.sliceB, got, c.want)
		}
	}
}

func TestInt64SliceEquals(t *testing.T) {
	testCases := []struct {
		sliceA []int64
		sliceB []int64
		want   bool
	}{
		{[]int64{}, []int64{}, true},
		{[]int64{1}, []int64{}, false},
		{[]int64{}, []int64{1}, false},
		{[]int64{1}, []int64{1}, true},
		{[]int64{1, 2}, []int64{1}, false},
		{[]int64{1, 2}, []int64{2}, false},
		{[]int64{1}, []int64{1, 2}, false},
		{[]int64{2}, []int64{1, 2}, false},
		{[]int64{1, 2}, []int64{1, 2}, true},
		{[]int64{1, 2}, []int64{1, 3}, false},
	}

	for _, c := range testCases {
		got := Int64SliceEquals(c.sliceA, c.sliceB)
		if got != c.want {
			t.Errorf("Int64SliceEquals(%v, %v) == %v, want %v", c.sliceA, c.sliceB, got, c.want)
		}
	}
}

func TestInt64SliceToStringSlice(t *testing.T) {
	testCases := []struct {
		intSlice []int64
		want     []string
	}{
		{[]int64{}, []string{}},
		{[]int64{3, 5, 6, 9}, []string{"3", "5", "6", "9"}},
	}

	for _, c := range testCases {
		got := Int64SliceToStringSlice(c.intSlice)
		if !commonstringutils.StringSliceCompare(got, c.want) {
			t.Errorf("Int64SliceToStringSlice(%d) == %s, want %s", c.intSlice, got, c.want)
		}
	}
}

func TestProductInt64Range(t *testing.T) {
	testCases := []struct {
		low  int64
		high int64
		want int64
	}{
		{0, 0, 0},
		{0, 1, 0},
		{1, 0, 0},
		{1, 2, 2},
		{1, 3, 6},
		{1, 4, 24},
	}

	for _, c := range testCases {
		got := ProductInt64Range(c.low, c.high)
		if got != c.want {
			t.Errorf("ProductInt64Slice(%v, %v) == %v, want %v", c.low, c.high, got, c.want)
		}
	}
}

func TestProductInt64Slice(t *testing.T) {
	testCases := []struct {
		intSlice []int64
		want     int64
	}{
		{[]int64{}, 0},
		{[]int64{0}, 0},
		{[]int64{1}, 1},
		{[]int64{1, 2}, 2},
		{[]int64{1, 2, 3}, 6},
	}

	for _, c := range testCases {
		got := ProductInt64Slice(c.intSlice)
		if got != c.want {
			t.Errorf("ProductInt64Slice(%v) == %v, want %v", c.intSlice, got, c.want)
		}
	}
}

func TestAtoInt64(t *testing.T) {
	testCases := []struct {
		s    string
		want int64
	}{
		{"0", 0},
		{"1", 1},
		{"10001", 10001},
	}

	for _, c := range testCases {
		got := AtoInt64(c.s)
		if got != c.want {
			t.Errorf("AtoInt64(%v) == %v, want %v", c.s, got, c.want)
		}
	}
}
