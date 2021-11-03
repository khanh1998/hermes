package utils

import (
	"log"
	"reflect"
	"testing"
)

type SearchTestResult struct {
	Index int
	Ok    bool
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
		Ints: []int{1, 2, 3},
		Cases: []SearchTestCase{
			{
				Search: 4,
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
				Search: 1,
				Result: SearchTestResult{
					Index: 0,
					Ok:    true,
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
				log.Fatalf(
					"%v search value:%v res index:%v found:%v",
					testSuite.Ints, testCase.Search, testCase.Result.Index, testCase.Result.Ok,
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
