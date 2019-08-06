package main

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
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

func TestSumWithTestify(t *testing.T) {
	a := assert.New(t)
	i := 1
	i2 := 2
	want := 3
	got := Sum(i, i2)
	a.EqualValues(want, got)
}

func TestSumWithCmp(t *testing.T) {
	i := 1
	i2 := 2
	want := 3
	got := Sum(i, i2)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Sum() mismatch (-want +got):\n%s", diff)
	}
}
