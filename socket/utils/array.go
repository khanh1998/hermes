package utils

import (
	"sort"
)

func SortAsc(ints []int) {
	sort.Ints(ints)
}

// BinarySearch search for a value in a sorted array.
// if it found the value, then start = stop = index of value in the array.
// if the value is not existed in the array, diff == 1,
// start and stop is the indexes of two values at the left and right of the searched value.
func Search(ints []int, value int) (index int, res int, ok bool) {
	length := len(ints)
	if length == 0 {
		return -1, -1, false
	}
	start, stop := 0, length-1
	diff := stop - start
	for diff != 0 {
		middle := int((start + stop) / 2)
		middleValue := ints[middle]
		if value < middleValue {
			stop = middle - 1
			if stop < start {
				stop = start
			}
		} else if value > middleValue {
			start = middle + 1
			if start >= length {
				start = length - 1
			}
		} else if value == middleValue {
			return middle, middleValue, true
		}
		diff = stop - start
	}
	searchedValue := ints[start]
	return start, searchedValue, searchedValue == value
}

// AddToSortedArr add a integer to a sorted array
func AddToSortedArr(ints []int, value int) []int {
	// now we looking for a position to insert new value
	index, searched, ok := Search(ints, value)
	// we actually insert value here
	if ok {
		// do nothing because the value is in the array
		return ints
	}
	if value < searched {
		tmp := make([]int, len(ints)+1)
		copy(tmp[0:], ints[:index])
		tmp[index] = value
		copy(tmp[index+1:], ints[index:])
		return tmp
	} else {
		tmp := make([]int, len(ints)+1)
		copy(tmp[0:], ints[:index+1])
		tmp[index+1] = value
		copy(tmp[index+2:], ints[index+1:])
		return tmp
	}
}

func RemoveFromSorted(ints []int, value int) []int {
	index, _, ok := Search(ints, value)
	if ok {
		return append(ints[:index], ints[index+1:]...)
	}
	return ints
}
