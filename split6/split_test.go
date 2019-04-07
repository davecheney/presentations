package split

import (
	"reflect"
	"testing"
)

// tag::test[]
func TestSplit(t *testing.T) {
	tests := []struct {
		name  string
		input string
		sep   string
		want  []string
	}{
		{name: "simple", input: "a/b/c", sep: "/", want: []string{"a", "b", "c"}},
		{name: "wrong sep", input: "a/b/c", sep: ",", want: []string{"a/b/c"}},
		{name: "no sep", input: "abc", sep: "/", want: []string{"abc"}},
		{name: "trailing sep", input: "a/b/c/", sep: "/", want: []string{"a", "b", "c"}},
	}

	for _, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}

// end::test[]
