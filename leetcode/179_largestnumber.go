package leetcode

import (
	"sort"
	"strconv"
	"strings"
)

type largestAfterConcatenation []string

func (nums largestAfterConcatenation) Len() int {
	return len(nums)
}

func (nums largestAfterConcatenation) Swap(i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func (nums largestAfterConcatenation) Less(i, j int) bool {
	return nums[i]+nums[j] > nums[j]+nums[i]
}

func largestNumber(nums []int) string {
	strs := []string{}
	for _, num := range nums {
		strs = append(strs, strconv.Itoa(num))
	}
	sort.Sort(largestAfterConcatenation(strs))
	if len(strs) == 0 || strs[0] == "0" {
		return "0"
	}
	return strings.Join(strs, "")
}
