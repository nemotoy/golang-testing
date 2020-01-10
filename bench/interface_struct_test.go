package main_test

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/matryer/is"
	"github.com/stretchr/testify/assert"
)

type object struct {
	S string
	I int
	B bool
}

func genObjectStub() (object, interface{}) {
	o := object{"s", 1, true}
	return o, o
}

func genPointerObjectStub() (*object, interface{}) {
	o := &object{"s", 1, true}
	return o, o
}

func BenchmarkInterfaceStructWithReflect(b *testing.B) {
	want, got := genObjectStub()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if !reflect.DeepEqual(got, want) {
			b.Fatal()
		}
	}
}

func BenchmarkInterfacePointerStructWithReflect(b *testing.B) {
	want, got := genPointerObjectStub()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if !reflect.DeepEqual(got, want) {
			b.Fatal()
		}
	}
}

func BenchmarkInterfaceStructWithTestify(b *testing.B) {
	want, got := genObjectStub()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a := assert.New(b)
		a.EqualValues(want, got)
	}
}

func BenchmarkInterfacePointerStructWithTestify(b *testing.B) {
	want, got := genPointerObjectStub()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a := assert.New(b)
		a.EqualValues(want, got)
	}
}

func BenchmarkInterfaceStructWithCmp(b *testing.B) {
	want, got := genObjectStub()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if !cmp.Equal(want, got) {
			b.Fatal()
		}
	}
}

func BenchmarkInterfacePointerStructWithCmp(b *testing.B) {
	want, got := genPointerObjectStub()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if !cmp.Equal(want, got) {
			b.Fatal()
		}
	}
}

func BenchmarkInterfaceStructWithIs(b *testing.B) {
	want, got := genObjectStub()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		is := is.New(b)
		is.Equal(want, got)
	}
}

func BenchmarkInterfacePointerStructWithIs(b *testing.B) {
	want, got := genPointerObjectStub()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		is := is.New(b)
		is.Equal(want, got)
	}
}
