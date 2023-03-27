package collection

import (
	"github.com/devsquared/gods/test"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestNewList(t *testing.T) {
	newList := NewList[int]()

	if newList.size != 0 {
		test.ReportTestFailure(t, newList.size, 0)
	}
}

func TestNewListFromSlice(t *testing.T) {
	startingSlice := []string{"ello there", "oy"}

	newList := NewListFromSlice[string](startingSlice)

	if !cmp.Equal(newList.coreSlice, startingSlice) {
		test.ReportTestFailure(t, newList.coreSlice, startingSlice)
	}

	if newList.size != len(startingSlice) {
		test.ReportTestFailure(t, newList.size, len(startingSlice))
	}
}

func TestList_Empty(t *testing.T) {
	startingSlice := []string{"ello there", "oy"}

	newList := NewListFromSlice[string](startingSlice)
	newList.Empty()

	expectedList := NewList[string]()

	actualCoreSlice := newList.coreSlice
	expectedCoreSlice := expectedList.coreSlice

	if !cmp.Equal(actualCoreSlice, expectedCoreSlice) {
		test.ReportTestFailure(t, actualCoreSlice, expectedCoreSlice)
	}

	actualSize := newList.size
	expectedSize := expectedList.size
	if actualSize != expectedSize {
		test.ReportTestFailure(t, actualSize, expectedSize)
	}
}

func TestList_Add(t *testing.T) {
	startingSlice := []string{"startingElement"}

	listFromSlice := NewListFromSlice[string](startingSlice)

	newElement := "newElement"
	expectedSlice := append(startingSlice, newElement)

	listFromSlice.Add(newElement)

	actualSlice := listFromSlice.coreSlice

	if !cmp.Equal(actualSlice, expectedSlice) {
		test.ReportTestFailure(t, actualSlice, expectedSlice)
	}
}
