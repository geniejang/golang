package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTreeNode(t *testing.T) {
	testcases := []struct {
		Val      string
		Expected *TreeNode
	}{
		{"", nil},
		{"null", nil},
		{"1", &TreeNode{Val: 1}},
		{"-7", &TreeNode{Val: -7}},
		{"0", &TreeNode{Val: 0}},
	}
	for _, tc := range testcases {
		assert.Equal(t, tc.Expected, newTreeNode(tc.Val))
	}
}

func TestNewTreeFromStringValues(t *testing.T) {
	testcases := []struct {
		StrArr   string
		Expected *TreeNode
	}{
		{"", nil},
		{"100", &TreeNode{Val: 100}},
		{"2,1,3", &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}},
		{
			"5,1,4,null,null,3,6",
			&TreeNode{
				Val:  5,
				Left: &TreeNode{Val: 1},
				Right: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 6},
				},
			},
		},
	}
	for _, tc := range testcases {
		assert.Equal(t, tc.Expected, newTreeFromStringValues(tc.StrArr), "arr: [%s]", tc.StrArr)
	}
}
