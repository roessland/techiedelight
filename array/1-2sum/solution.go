package main

import (
	"errors"
	"fmt"
	"sort"
)

// Find two numbers in an array that sum to a specific number and returns their
// indicies. The indices are always different if the result is valid.
//
// nlogn: Sorts array and uses binary search
func TwoSum(unsorted []int, sum int) (int, int, error) {
	arr := make([]int, len(unsorted))
	copy(arr, unsorted)
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		j := sort.SearchInts(arr, sum-arr[i])
		if j == len(arr) {
			continue
		}
		if j == i {
			continue
		}
		if arr[i]+arr[j] == sum {
			a := indexOf(unsorted, arr[i], 0)
			b := indexOf(unsorted, arr[j], 0)
			if a == b {
				b = indexOf(unsorted, arr[j], a+1)
			}
			return a, b, nil
		}
	}
	return -1, -1, errors.New("no such pair")
}

func indexOf(arr []int, n int, start int) int {
	for i := start; i < len(arr); i++ {
		if arr[i] == n {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(TwoSum([]int{8, 7, 2, 5, 3, 1}, 10))
}
