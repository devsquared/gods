package test

import "fmt"

func ReportTestFailure(scenarioName string, got, wanted any) string {
	return fmt.Sprintf("scenario: %s \n\t got: %v, wanted: %v", scenarioName, got, wanted)
}

func CheckErrorsAreSame(actualErr, expectedErr error) bool {
	if actualErr != nil && expectedErr != nil {
		if actualErr.Error() == expectedErr.Error() {
			// if both errors are non-nil and the same, they are the same
			return true
		}
	}

	return false
}
