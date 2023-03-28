package test

import (
	"fmt"
	"testing"
)

func ReportTestFailure(t *testing.T, got, wanted any) {
	t.Errorf(fmt.Sprintf("scenario: %s \n\t got: %v, wanted: %v", t.Name(), got, wanted))
}

func IsErrSame(actualErr, expectedErr error) bool {
	if actualErr == nil && expectedErr == nil {
		return true
	}

	if actualErr != nil && expectedErr != nil {
		if actualErr.Error() == expectedErr.Error() {
			// if both errors are non-nil and the same, they are the same
			return true
		}
	}

	return false
}
