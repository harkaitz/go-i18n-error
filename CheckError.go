package i18ne

import (
	"strings"
	"testing"
)

// CheckError checkes whether the error is not nil and contains the
// code between parentheses.
func CheckError(t *testing.T, e error, code string) error {
	if e == nil {
		t.Fatalf("Should have failed with code (%v)", code)
	}
	if !strings.Contains(e.Error(), "(" + code + ")") {
		t.Fatalf("Should have failed with %v: %s", code, e)
	}
	return nil
}

// CheckErrors is the same as CheckError but for a slice of errors.
func CheckErrors(t *testing.T, el []error, code string) error {
	return CheckError(t, JoinErrors(el), code)
}
