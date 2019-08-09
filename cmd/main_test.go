package main_test

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/matryer/is"
	"github.com/stretchr/testify/assert"
)

func Sum(i, i2 int) int {
	return i + i2
}

func BenchmarkSumWithReflect(b *testing.B) {
	i, i2 := 1, 2
	want := i + i2
	got := Sum(i, i2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if !reflect.DeepEqual(got, want) {
			b.Fatal()
		}
	}
}

func BenchmarkSumWithTestify(b *testing.B) {
	i, i2 := 1, 2
	want := i + i2
	got := Sum(i, i2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a := assert.New(b)
		a.EqualValues(want, got)
	}
}

func BenchmarkSumWithCmp(b *testing.B) {
	i, i2 := 1, 2
	want := i + i2
	got := Sum(i, i2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if diff := cmp.Diff(want, got); diff != "" {
			b.Errorf("Sum() mismatch (-want +got):\n%s", diff)
		}
	}
}

func BenchmarkSumWithIs(b *testing.B) {
	i, i2 := 1, 2
	want := i + i2
	got := Sum(i, i2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		is := is.New(b)
		is.Equal(want, got)
	}
}
