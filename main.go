package assert

import (
	"reflect"
	"testing"
)

// toggles if it should announce succeeding cases
var HideSuccess = false

// Assert makes sure value is true, exits with custom error message
func Assert(t *testing.T, val bool, failReason string) {
	t.Helper()

	if !val {
		t.Fatalf("fail: %s\n", failReason)
	}
}

// Equals makes sure values are equal
func Equals(t *testing.T, a any, b any, customErrorMessage ...string) {
	t.Helper()

	if !reflect.DeepEqual(a, b) {
		if len(customErrorMessage) > 0 {
			t.Fatalf("fail: %s\n", customErrorMessage[0])
		} else {
			t.Fatalf("fail: '%v' is not '%v'\n", a, b)
		}
	} else if !HideSuccess {
		t.Logf("success: '%v' is '%v'\n", a, b)
	}
}

// NotEquals makes sure values arent equal
func NotEquals(t *testing.T, a any, b any, customErrorMessage ...string) {
	t.Helper()

	if !reflect.DeepEqual(a, b) {
		if !HideSuccess {
			t.Logf("success: '%v' is not '%v'\n", a, b)
		}
	} else {
		if len(customErrorMessage) > 0 {
			t.Fatalf("fail: %s\n", customErrorMessage[0])
		} else {
			t.Fatalf("fail: '%v' is '%v'\n", a, b)
		}
	}
}
