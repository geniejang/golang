package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidBST(t *testing.T) {
	testcases := []struct {
		Root     *TreeNode
		Expected bool
	}{
		{newTreeFromStringValues("2,1,3"), true},
		{newTreeFromStringValues("5,1,4,null,null,3,6"), false},
		{nil, true},
		{newTreeNode("1"), true},
		{newTreeFromStringValues("1,2"), false},
		{newTreeFromStringValues("3,null,1"), false},
	}
	for _, tc := range testcases {
		assert.Equal(t, tc.Expected, isValidBST(tc.Root), tc.Root.JsonString())
	}
}
