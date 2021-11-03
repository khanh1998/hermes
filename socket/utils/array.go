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
		} else if value > middleValue {
			start = middle + 1
		} else if value == middleValue {
			break
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
		new := append(ints, 0)
		copy(new[:index+2], new[:index+1])
		new[index+1] = value
		return new
	}
}

func RemoveFromSorted(ints []int, value int) []int {
	index, _, ok := Search(ints, value)
	if ok {
		return append(ints[:index], ints[index+1:]...)
	}
	return ints
}
