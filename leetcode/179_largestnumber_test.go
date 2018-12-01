package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertLargestNumber(t *testing.T, nums []int, expected string) {
	actual := largestNumber(nums)
	assert.Equal(t, expected, actual)
}

func TestLargestNumber_GivenExample1(t *testing.T) {
	nums := []int{10, 2}
	expected := "210"
	assertLargestNumber(t, nums, expected)
}

func TestLargestNumber_GivenExample2(t *testing.T) {
	nums := []int{3, 30, 34, 5, 9}
	expected := "9534330"
	assertLargestNumber(t, nums, expected)
}

func TestLargestNumber_EmptyInput(t *testing.T) {
	nums := []int{}
	expected := "0"
	assertLargestNumber(t, nums, expected)
}

func TestLargestNumber_MultipleZeroesReturnsSingleZero(t *testing.T) {
	nums := []int{0, 0}
	expected := "0"
	assertLargestNumber(t, nums, expected)
}
