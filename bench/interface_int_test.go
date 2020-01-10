package main_test

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/matryer/is"
	"github.com/stretchr/testify/assert"
)

func getSumInterface(i, i2 int) interface{} {
	return i + i2
}

func genStub() (int, interface{}) {
	i, i2 := 1, 2
	want := i + i2
	return want, getSumInterface(i, i2)
}

func BenchmarkInterfaceIntWithReflect(b *testing.B) {
	want, got := genStub()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if !reflect.DeepEqual(got, want) {
			b.Fatal()
		}
	}
}

func BenchmarkInterfaceIntWithTestify(b *testing.B) {
	want, got := genStub()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a := assert.New(b)
		a.EqualValues(want, got)
	}
}

func BenchmarkInterfaceIntWithCmp(b *testing.B) {
	want, got := genStub()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if !cmp.Equal(want, got) {
			b.Fatal()
		}
	}
}

func BenchmarkInterfaceIntWithIs(b *testing.B) {
	want, got := genStub()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		is := is.New(b)
		is.Equal(want, got)
	}
}
