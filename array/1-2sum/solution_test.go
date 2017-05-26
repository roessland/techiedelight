package main

import "testing"

import "github.com/stretchr/testify/assert"

func TestZeroLength(t *testing.T) {
	i0, i1, err := TwoSum([]int{}, 10)
	assert.Equal(t, -1, i0)
	assert.Equal(t, -1, i1)
	assert.NotNil(t, err)
}

func TestEasy(t *testing.T) {
	arr := []int{1, 2}
	sum := 3
	i, j, err := TwoSum(arr, sum)
	assert.Equal(t, sum, arr[i]+arr[j])
	assert.NotEqual(t, i, j)
	assert.Nil(t, err)
}

func TestSmall(t *testing.T) {
	arr := []int{8, 7, 2, 5, 3, 1}
	sum := 10
	i, j, err := TwoSum(arr, sum)
	assert.Equal(t, sum, arr[i]+arr[j])
	assert.NotEqual(t, i, j)
	assert.Nil(t, err)
}

func TestNotFound(t *testing.T) {
	arr := []int{10, 15, 20, 30}
	sum := 15
	i, j, err := TwoSum(arr, sum)
	assert.Equal(t, -1, i)
	assert.Equal(t, -1, j)
	assert.NotNil(t, err)
}

func TestSameNumber(t *testing.T) {
	arr := []int{10, 10}
	sum := 20
	i, j, err := TwoSum(arr, sum)
	assert.Equal(t, sum, arr[i]+arr[j])
	assert.NotEqual(t, i, j)
	assert.Nil(t, err)
}
