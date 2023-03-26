package collection

import (
	"github.com/devsquared/gods/test"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_NewList(t *testing.T) {
	newList := NewList[int]()

	if newList.size != 0 {
		test.ReportTestFailure(t.Name(), newList.size, 0)
	}
}

func Test_NewListFromSlice(t *testing.T) {
	startingSlice := []string{"ello there", "oy"}

	newList := NewListFromSlice[string](startingSlice)

	if !cmp.Equal(newList.coreSlice, startingSlice) {
		test.ReportTestFailure(t.Name(), newList.coreSlice, startingSlice)
	}

	if newList.size != len(startingSlice) {
		test.ReportTestFailure(t.Name(), newList.size, len(startingSlice))
	}
}

func Test_Empty(t *testing.T) {
	startingSlice := []string{"ello there", "oy"}

	newList := NewListFromSlice[string](startingSlice)
	newList.Empty()

	expectedList := NewList[string]()

	actualCoreSlice := newList.coreSlice
	expectedCoreSlice := expectedList.coreSlice

	if !cmp.Equal(actualCoreSlice, expectedCoreSlice) {
		test.ReportTestFailure(t.Name(), actualCoreSlice, expectedCoreSlice)
	}

	actualSize := newList.size
	expectedSize := expectedList.size
	if actualSize != expectedSize {
		test.ReportTestFailure(t.Name(), actualSize, expectedSize)
	}
}
