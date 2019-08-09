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
	for i := 0; i < b.N; i++ {
		i, i2 := i, i+1
		want := i + i2
		got := Sum(i, i2)
		if !reflect.DeepEqual(got, want) {
			b.Fatal()
		}
	}
}

func BenchmarkSumWithTestify(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := assert.New(b)
		i, i2 := i, i+1
		want := i + i2
		got := Sum(i, i2)
		a.EqualValues(want, got)
	}
}

func BenchmarkSumWithCmp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i, i2 := i, i+1
		want := i + i2
		got := Sum(i, i2)
		if diff := cmp.Diff(want, got); diff != "" {
			b.Errorf("Sum() mismatch (-want +got):\n%s", diff)
		}
	}
}

func BenchmarkSumWithIs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		is := is.New(b)
		i, i2 := i, i+1
		want := i + i2
		got := Sum(i, i2)
		is.Equal(want, got)
	}
}
