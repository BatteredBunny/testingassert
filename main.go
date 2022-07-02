package testingassert

import (
	"reflect"
	"testing"
)

// Assert / makes sure value is true, exists with custom error message
func Assert(t *testing.T, val bool, message string) {
	t.Helper()

	if !val {
		t.Fatal(message)
	}
}

// AssertEquals / Compares comparable types
func AssertEquals[T comparable](t *testing.T, first T, second T) {
	t.Helper()

	if first != second {
		t.Fatalf("'%v' is not '%v'\n", first, second)
	} else {
		t.Logf("'%v' is '%v'\n", first, second)
	}
}

// AssertEqualsMsg / AssertEquals with custom error message
func AssertEqualsMsg[T comparable](t *testing.T, first T, second T, message string) {
	t.Helper()

	if first != second {
		t.Fatal(message)
	} else {
		t.Logf("'%v' is '%v'\n", first, second)
	}
}

// AssertEqualsDeep / AssertEquals with reflect
func AssertEqualsDeep[T any](t *testing.T, first T, second T) {
	t.Helper()

	if !reflect.DeepEqual(first, second) {
		t.Fatalf("'%v' is not '%v'\n", first, second)
	} else {
		t.Logf("'%v' is '%v'\n", first, second)
	}
}

// AssertEqualsDeepMsg / AssertEqualsDeep with custom error message
func AssertEqualsDeepMsg[T any](t *testing.T, first T, second T, message string) {
	t.Helper()

	if !reflect.DeepEqual(first, second) {
		t.Fatal(message)
	} else {
		t.Logf("'%v' is '%v'\n", first, second)
	}
}
