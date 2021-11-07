package utils

import (
	"log"
	"reflect"
	"testing"
)

type SearchTestResult struct {
	Index int  // nearest number index in case not found
	Ok    bool // found ?
}
type SearchTestCase struct {
	Search int
	Result SearchTestResult
}
type SearchTestSuite struct {
	Ints  []int
	Cases []SearchTestCase
}

var testSuites []SearchTestSuite = []SearchTestSuite{
	{
		Ints: []int{},
		Cases: []SearchTestCase{
			{
				Search: 1,
				Result: SearchTestResult{
					Index: -1,
					Ok:    false,
				},
			},
		},
	},
	{
		Ints: []int{5},
		Cases: []SearchTestCase{
			{
				Search: 5,
				Result: SearchTestResult{
					Index: 0,
					Ok:    true,
				},
			},
			{
				Search: 1,
				Result: SearchTestResult{
					Index: 0,
					Ok:    false,
				},
			},
			{
				Search: 7,
				Result: SearchTestResult{
					Index: 0,
					Ok:    false,
				},
			},
		},
	},
	{
		Ints: []int{4, 6},
		Cases: []SearchTestCase{
			{
				Search: 4,
				Result: SearchTestResult{
					Index: 0,
					Ok:    true,
				},
			},
			{
				Search: 6,
				Result: SearchTestResult{
					Index: 1,
					Ok:    true,
				},
			},
			{
				Search: 1,
				Result: SearchTestResult{
					Index: 0,
					Ok:    false,
				},
			},
			{
				Search: 9,
				Result: SearchTestResult{
					Index: 1,
					Ok:    false,
				},
			},
		},
	},
	{
		Ints: []int{1, 5, 9},
		Cases: []SearchTestCase{
			{
				Search: 12,
				Result: SearchTestResult{
					Index: 2,
					Ok:    false,
				},
			},
			{
				Search: 4,
				Result: SearchTestResult{
					Index: 0,
					Ok:    false,
				},
			},
			{
				Search: 7,
				Result: SearchTestResult{
					Index: 2,
					Ok:    false,
				},
			},
			{
				Search: 0,
				Result: SearchTestResult{
					Index: 0,
					Ok:    false,
				},
			},
			{
				Search: -5,
				Result: SearchTestResult{
					Index: 0,
					Ok:    false,
				},
			},
			{
				Search: 1,
				Result: SearchTestResult{
					Index: 0,
					Ok:    true,
				},
			},
			{
				Search: 5,
				Result: SearchTestResult{
					Index: 1,
					Ok:    true,
				},
			},
			{
				Search: 9,
				Result: SearchTestResult{
					Index: 2,
					Ok:    true,
				},
			},
		},
	},
	{
		Ints: []int{0, 4, 8, 12},
		Cases: []SearchTestCase{
			{
				Search: 0,
				Result: SearchTestResult{
					Index: 0,
					Ok:    true,
				},
			},
			{
				Search: 4,
				Result: SearchTestResult{
					Index: 1,
					Ok:    true,
				},
			},
			{
				Search: 12,
				Result: SearchTestResult{
					Index: 3,
					Ok:    true,
				},
			},
			{
				Search: -1,
				Result: SearchTestResult{
					Index: 0,
					Ok:    false,
				},
			},
			{
				Search: 2,
				Result: SearchTestResult{
					Index: 0,
					Ok:    false,
				},
			},
			{
				Search: 7,
				Result: SearchTestResult{
					Index: 2,
					Ok:    false,
				},
			},
			{
				Search: 9,
				Result: SearchTestResult{
					Index: 3,
					Ok:    false,
				},
			},
			{
				Search: 13,
				Result: SearchTestResult{
					Index: 3,
					Ok:    false,
				},
			},
		},
	},
}

func TestSearch(t *testing.T) {
	for index, testSuite := range testSuites {
		t.Logf("Suite %v ------------------- %v", index, testSuite.Ints)
		for _, testCase := range testSuite.Cases {
			index, _, ok := Search(testSuite.Ints, testCase.Search)
			if index == testCase.Result.Index && ok == testCase.Result.Ok {
				continue
			} else {
				t.Logf(
					"input: %v search value:%v res index:%v found:%v",
					testSuite.Ints, testCase.Search, testCase.Result.Index, testCase.Result.Ok,
				)
				t.Fatalf(
					"result: index:%v found:%v",
					index, ok,
				)
			}
		}
	}
}

type AddToArrTestCase struct {
	Add    int
	Result []int
}
type AddToArrTestSuite struct {
	Ints  []int
	Cases []AddToArrTestCase
}

var addingTestSuites []AddToArrTestSuite = []AddToArrTestSuite{
	{
		Ints: []int{},
		Cases: []AddToArrTestCase{
			{
				Add:    1,
				Result: []int{1},
			},
			{
				Add:    2,
				Result: []int{2},
			},
		},
	},
	{
		Ints: []int{1},
		Cases: []AddToArrTestCase{
			{
				Add:    2,
				Result: []int{1, 2},
			},
			{
				Add:    0,
				Result: []int{0, 1},
			},
			{
				Add:    1,
				Result: []int{1},
			},
			{
				Add:    0,
				Result: []int{0, 1},
			},
		},
	},
	{
		Ints: []int{1, 3, 5},
		Cases: []AddToArrTestCase{
			{
				Add:    4,
				Result: []int{1, 3, 4, 5},
			},
			{
				Add:    6,
				Result: []int{1, 3, 5, 6},
			},
			{
				Add:    2,
				Result: []int{1, 2, 3, 5},
			},
			{
				Add:    0,
				Result: []int{0, 1, 3, 5},
			},
			{
				Add:    1,
				Result: []int{1, 3, 5},
			},
			{
				Add:    3,
				Result: []int{1, 3, 5},
			},
			{
				Add:    5,
				Result: []int{1, 3, 5},
			},
		},
	},
	{
		Ints: []int{2, 4, 6, 8},
		Cases: []AddToArrTestCase{
			{
				Add:    11,
				Result: []int{2, 4, 6, 8, 11},
			},
			{
				Add:    7,
				Result: []int{2, 4, 6, 7, 8},
			},
			{
				Add:    5,
				Result: []int{2, 4, 5, 6, 8},
			},
			{
				Add:    3,
				Result: []int{2, 3, 4, 6, 8},
			},
			{
				Add:    -1,
				Result: []int{-1, 2, 4, 6, 8},
			},
		},
	},
}

func TestAddToSortedArr(t *testing.T) {
	for index, testSuite := range addingTestSuites {
		t.Logf("Test suite: %v -------------------- %v", index, testSuite.Ints)
		for _, testCase := range testSuite.Cases {
			input := make([]int, len(testSuite.Ints))
			copy(input, testSuite.Ints)
			res := AddToSortedArr(input, testCase.Add)
			if same := reflect.DeepEqual(res, testCase.Result); !same {
				t.Logf("output: %v", res)
				t.Fatalf("%v add:%v res:%v\n", testSuite.Ints, testCase.Add, testCase.Result)
			}
		}
	}
}

type RemoveFromArrCase struct {
	Remove int
	Result []int
}
type RemoveFromArrSuite struct {
	Ints  []int
	Cases []RemoveFromArrCase
}

var removeFromArrSuites []RemoveFromArrSuite = []RemoveFromArrSuite{
	{
		Ints: []int{},
		Cases: []RemoveFromArrCase{
			{
				Remove: 4,
				Result: []int{},
			},
		},
	},
	{
		Ints: []int{5},
		Cases: []RemoveFromArrCase{
			{
				Remove: 1,
				Result: []int{5},
			},
			{
				Remove: 8,
				Result: []int{5},
			},
			{
				Remove: 5,
				Result: []int{},
			},
		},
	},
	{
		Ints: []int{1, 7},
		Cases: []RemoveFromArrCase{
			{
				Remove: 1,
				Result: []int{7},
			},
			{
				Remove: 7,
				Result: []int{1},
			},
			{
				Remove: 0,
				Result: []int{1, 7},
			},
			{
				Remove: 5,
				Result: []int{1, 7},
			},
			{
				Remove: 9,
				Result: []int{1, 7},
			},
		},
	},
	{
		Ints: []int{1, 5, 9},
		Cases: []RemoveFromArrCase{
			{
				Remove: 1,
				Result: []int{5, 9},
			},
			{
				Remove: 5,
				Result: []int{1, 9},
			},
			{
				Remove: 9,
				Result: []int{1, 5},
			},
			{
				Remove: 10,
				Result: []int{1, 5, 9},
			},
			{
				Remove: 7,
				Result: []int{1, 5, 9},
			},
			{
				Remove: 3,
				Result: []int{1, 5, 9},
			},
			{
				Remove: -10,
				Result: []int{1, 5, 9},
			},
		},
	},
}

func TestRemoveFromSortedArr(t *testing.T) {
	for index, testSuite := range removeFromArrSuites {
		t.Logf("Test suite: %v ------------------- %v", index, testSuite.Ints)
		for _, testCase := range testSuite.Cases {
			input := make([]int, len(testSuite.Ints))
			copy(input, testSuite.Ints)
			res := RemoveFromSorted(input, testCase.Remove)
			if same := reflect.DeepEqual(res, testCase.Result); !same {
				log.Printf("output: %v", res)
				log.Fatalf("ints: %v remove: %v expected: %v", testSuite.Ints, testCase.Remove, testCase.Result)
			}
		}
	}
}
