package main

import "testing"

import "github.com/stretchr/testify/assert"

func TestZeroLength(t *testing.T) {
	i0, i1, err := TwoSum([]int{}, 10)
	assert.NotNil(t, err)
	assert.Equal(t, -1, i0)
	assert.Equal(t, -1, i1)
}
