package main

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/matryer/is"
	"github.com/stretchr/testify/assert"
)

func TestSumWithReflect(t *testing.T) {
	i, i2 := 1, 2
	want := 3
	got := Sum(i, i2)
	if !reflect.DeepEqual(got, want) {
		t.Fatal()
	}
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

func TestSumWithTestify(t *testing.T) {
	a := assert.New(t)
	i, i2 := 1, 2
	want := 3
	got := Sum(i, i2)
	a.EqualValues(want, got)
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

func TestSumWithCmp(t *testing.T) {
	i, i2 := 1, 2
	want := 3
	got := Sum(i, i2)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Sum() mismatch (-want +got):\n%s", diff)
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

func TestSumWithIs(t *testing.T) {
	is := is.New(t)
	i, i2 := 1, 2
	want := 3
	got := Sum(i, i2)
	is.Equal(want, got)
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

func TestObjectWithCmp(t *testing.T) {
	want := Object{"a", 1, true}
	got := genObject("a", 1, true)
	if !cmp.Equal(want, got) {
		t.Error("genObject() doesn't equal")
	}
}
