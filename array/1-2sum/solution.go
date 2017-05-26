package main

import (
	"errors"
	"fmt"
	"sort"
)

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
			return indexOf(unsorted, arr[i]), indexOf(unsorted, arr[j]), nil
		}
	}
	return -1, -1, errors.New("no such pair")
}

func indexOf(arr []int, n int) int {
	for i, val := range arr {
		if val == n {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(TwoSum([]int{8, 7, 2, 5, 3, 1}, 10))
}
