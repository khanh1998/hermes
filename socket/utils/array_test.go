package utils

import (
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
		t.Logf("Suite %v -------------------", index)
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
		Ints: []int{1, 3, 5},
		Cases: []AddToArrTestCase{
			{
				Add:    4,
				Result: []int{1, 3, 4, 5},
			},
		},
	},
}

func TestAddToSortedArr(t *testing.T) {
	for index, testSuite := range addingTestSuites {
		t.Logf("Test suite: %v --------------------", index)
		for _, testCase := range testSuite.Cases {
			input := make([]int, len(testSuite.Ints))
			copy(input, testSuite.Ints)
			res := AddToSortedArr(input, testCase.Add)
			if same := reflect.DeepEqual(res, testCase.Result); !same {
				t.Fatalf("%v add:%v res:%v\n", testSuite.Ints, testCase.Add, testCase.Result)
			}
		}
	}
}
