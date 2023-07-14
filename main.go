package assert

import (
	"reflect"
	"testing"
)

// toggles if it should announce succeeding cases
var HideSuccess = false

// the testing function state
var TestState *testing.T

// Assert makes sure value is true, exits with custom error message
func Assert(val bool, failReason string) {
	TestState.Helper()

	if !val {
		TestState.Fatalf("fail: %s\n", failReason)
	}
}

// Equals makes sure values are equal
func Equals(a any, b any, customErrorMessage ...string) {
	TestState.Helper()

	if !reflect.DeepEqual(a, b) {
		if len(customErrorMessage) > 0 {
			TestState.Fatalf("fail: %s\n", customErrorMessage[0])
		} else {
			TestState.Fatalf("fail: '%v' is not '%v'\n", a, b)
		}
	} else if !HideSuccess {
		TestState.Logf("success: '%v' is '%v'\n", a, b)
	}
}

// NotEquals makes sure values arent equal
func NotEquals(a any, b any, customErrorMessage ...string) {
	TestState.Helper()

	if !reflect.DeepEqual(a, b) {
		if !HideSuccess {
			TestState.Logf("success: '%v' is not '%v'\n", a, b)
		}
	} else {
		if len(customErrorMessage) > 0 {
			TestState.Fatalf("fail: %s\n", customErrorMessage[0])
		} else {
			TestState.Fatalf("fail: '%v' is '%v'\n", a, b)
		}
	}
}
