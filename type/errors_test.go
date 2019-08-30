package errors

import (
	"fmt"
	"reflect"
	"testing"
)

type object struct {
	err error
}

func (o object) Error() string {
	return fmt.Sprintf("ERROR: %v", o.err)
}

func TestGenObject(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   int
		wb   bool
		err  error
		ws   string
	}{
		{"#1", 1, false, nil, ""},
		{"#2", 2, true, fmt.Errorf("ERROR: 2"), "ERROR: 2"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := genObject(tc.in)
			if tc.wb && got == nil {
				t.Fail()
			}
			if tc.wb && !reflect.DeepEqual(got.err, tc.err) {
				t.Errorf("want: %v, got: %v", tc.err, got)
			}
			if tc.wb && tc.ws != got.err.Error() {
				t.Errorf("want: %v, got: %v", tc.ws, got.err.Error())
			}
		})
	}
}

func genObject(i int) *object {
	if i%2 == 0 {
		return &object{err: fmt.Errorf("ERROR: %v", i)}
	}
	return nil
}
