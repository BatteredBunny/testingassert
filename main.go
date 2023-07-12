package testingassert

import (
	"reflect"
	"testing"
)

// Assert makes sure value is true, exits with custom error message
func Assert(t *testing.T, val bool, message string) {
	t.Helper()

	if !val {
		t.Fatal(message)
	}
}

// AssertEquals compares comparable types
func AssertEquals[T comparable](t *testing.T, first T, second T) {
	t.Helper()

	if first != second {
		t.Fatalf("'%v' is not '%v'\n", first, second)
	} else {
		t.Logf("'%v' is '%v'\n", first, second)
	}
}

// AssertEqualsMsg is the error message variant of AssertEquals
func AssertEqualsMsg[T comparable](t *testing.T, first T, second T, message string) {
	t.Helper()

	if first != second {
		t.Fatal(message)
	} else {
		t.Logf("'%v' is '%v'\n", first, second)
	}
}

// AssertEqualsDeep is AssertEquals with reflect
func AssertEqualsDeep(t *testing.T, first any, second any) {
	t.Helper()

	if !reflect.DeepEqual(first, second) {
		t.Fatalf("'%v' is not '%v'\n", first, second)
	} else {
		t.Logf("'%v' is '%v'\n", first, second)
	}
}

// AssertEqualsDeepMsg is the error message variant of AssertEqualsDeep
func AssertEqualsDeepMsg(t *testing.T, first any, second any, message string) {
	t.Helper()

	if !reflect.DeepEqual(first, second) {
		t.Fatal(message)
	} else {
		t.Logf("'%v' is '%v'\n", first, second)
	}
}
