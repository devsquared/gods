package collection

import (
	"github.com/devsquared/gods/test"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestNewList(t *testing.T) {
	type testScenario struct {
		name          string
		expectedSlice []int
	}

	testScenarios := []testScenario{
		{
			name:          "basic new list",
			expectedSlice: []int{},
		},
	}

	for _, ts := range testScenarios {
		t.Run(ts.name, func(t *testing.T) {
			newList := NewList[int]()

			if !cmp.Equal(newList.coreSlice, ts.expectedSlice) {
				test.ReportTestFailure(t, newList.coreSlice, ts.expectedSlice)
			}
		})
	}

}

func TestNewListFromSlice(t *testing.T) {
	type testScenario struct {
		name          string
		inputSlice    []string
		expectedSlice []string
		expectedSize  int
	}

	testScenarios := []testScenario{
		{
			name:          "construct from empty slice",
			inputSlice:    []string{},
			expectedSlice: []string{},
		},
		{
			name:          "construct from non-empty slice",
			inputSlice:    []string{"ello there", "oy"},
			expectedSlice: []string{"ello there", "oy"},
		},
	}

	for _, ts := range testScenarios {
		t.Run(ts.name, func(t *testing.T) {
			newList := NewListFromSlice[string](ts.inputSlice)

			if !cmp.Equal(newList.coreSlice, ts.expectedSlice) {
				test.ReportTestFailure(t, newList.coreSlice, ts.expectedSlice)
			}
		})
	}
}

func TestList_Empty(t *testing.T) {
	type testScenario struct {
		name          string
		startingSlice []bool
	}

	testScenarios := []testScenario{
		{
			name:          "empty an empty list",
			startingSlice: []bool{},
		},
		{
			name:          "empty a non-empty list",
			startingSlice: []bool{true, false, true},
		},
	}

	for _, ts := range testScenarios {
		t.Run(ts.name, func(t *testing.T) {
			newList := NewListFromSlice[bool](ts.startingSlice)
			newList.Empty()

			expectedSlice := make([]bool, 0)

			if !cmp.Equal(expectedSlice, newList.coreSlice) {
				test.ReportTestFailure(t, newList.coreSlice, expectedSlice)
			}
		})
	}
}

func TestList_Add(t *testing.T) {
	type testScenario struct {
		name          string
		input         string
		expectedSlice []string
	}

	testScenarios := []testScenario{
		{
			name:          "add something",
			input:         "newElement",
			expectedSlice: []string{"startingElement", "newElement"},
		},
	}

	for _, ts := range testScenarios {
		t.Run(ts.name, func(t *testing.T) {
			startingSlice := []string{"startingElement"}
			listFromSlice := NewListFromSlice[string](startingSlice)

			listFromSlice.Add(ts.input)

			actualSlice := listFromSlice.coreSlice

			if !cmp.Equal(actualSlice, ts.expectedSlice) {
				test.ReportTestFailure(t, actualSlice, ts.expectedSlice)
			}
		})
	}
}

func TestList_RemovePanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected List.Set to panic with index out of bounds")
		}
	}()

	list := NewList[int]()
	list.Remove(-1)
}

func TestList_Remove(t *testing.T) {
	type testScenario struct {
		name          string
		startingSlice []string
		indexToRemove int
		expectedSlice []string
	}

	testScenarios := []testScenario{
		{
			name:          "remove from list",
			startingSlice: []string{"element", "toRemove"},
			indexToRemove: 1,
			expectedSlice: []string{"element"},
		},
	}

	for _, ts := range testScenarios {
		t.Run(ts.name, func(t *testing.T) {
			listFromSLice := NewListFromSlice[string](ts.startingSlice)

			listFromSLice.Remove(1)
			actualRemainingSlice := listFromSLice.coreSlice

			if !cmp.Equal(ts.expectedSlice, actualRemainingSlice) {
				test.ReportTestFailure(t, actualRemainingSlice, ts.expectedSlice)
			}
		})
	}
}

func TestList_SetPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected List.Set to panic with index out of bounds")
		}
	}()

	list := NewList[int]()
	list.Set(3, -1)
}

func TestList_Set(t *testing.T) {
	type testScenario struct {
		name          string
		input         string
		indexToSet    int
		expectedSlice []string
	}

	testScenarios := []testScenario{
		{
			name:          "add an element",
			input:         "test",
			indexToSet:    0,
			expectedSlice: []string{"test"},
		},
	}

	for _, scenario := range testScenarios {
		t.Run(scenario.name, func(t *testing.T) {
			list := NewList[string]()
			list.Set(scenario.input, scenario.indexToSet)

			if !cmp.Equal(list.coreSlice, scenario.expectedSlice) {
				test.ReportTestFailure(t, list.coreSlice, scenario.expectedSlice)
			}
		})
	}
}

func TestList_GetPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected List.Get to panic with index out of bounds")
		}
	}()

	list := NewList[int]()
	list.Get(-1)
}

func TestList_Get(t *testing.T) {
	type testScenario struct {
		name          string
		indexToGet    int
		expectedValue string
	}

	testScenarios := []testScenario{
		{
			name:          "get within bounds",
			indexToGet:    0,
			expectedValue: "test",
		},
	}

	for _, ts := range testScenarios {
		t.Run(ts.name, func(t *testing.T) {
			newList := NewList[string]()
			newList.Add(ts.expectedValue)

			actualValue := newList.Get(0)

			if !cmp.Equal(ts.expectedValue, actualValue) {
				test.ReportTestFailure(t, actualValue, ts.expectedValue)
			}
		})
	}

}
