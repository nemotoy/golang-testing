package main

import (
	"reflect"
	"testing"
)

func TestSumWithReflect(t *testing.T) {
	i := 1
	i2 := 2
	want := 3
	got := Sum(i, i2)
	if !reflect.DeepEqual(got, want) {
		t.Fatal()
	}
}
