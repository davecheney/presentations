// +build none

// tag::test[]
package split

import (
	"reflect"
	"testing"
)

// tag::tests[]
func TestSplit(t *testing.T) {
	got := Split("a/b/c", "/")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

// end::test[]

// tag::test2[]
func TestSplitWrongSep(t *testing.T) {
	got := Split("a/b/c", ",")
	want := []string{"a/b/c"} // <1>
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

// end::test2[]

// tag::empty[]
func testSplitEmptySep(t *testing.T) {
	got := Split("a/b/c", "")
	want := []string{"a/b/c"}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

// end::empty[]

// tag::test3[]
func testSplitTrailingSep(t *testing.T) {
	got := Split("a/b/c/", "/")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

// end::test3[]
// end::tests[]
