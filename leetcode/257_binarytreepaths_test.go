package leetcode

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertPaths(t *testing.T, expected, actual []string) {
	sort.Strings(expected)
	sort.Strings(actual)
	assert.Equal(t, expected, actual)
}

func TestBinaryTreePaths_GivenExample(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	expected := []string{"1->2->5", "1->3"}
	actual := binaryTreePaths(tree)
	assertPaths(t, expected, actual)
}

func TestBinaryTreePaths_NilReturnsEmptySlice(t *testing.T) {
	actual := binaryTreePaths(nil)
	assert.NotNil(t, actual)
	assert.Len(t, actual, 0)
}

func TestBinaryTreePaths_SingleNode(t *testing.T) {
	tree := &TreeNode{Val: 1, Left: nil, Right: nil}
	expected := []string{"1"}
	actual := binaryTreePaths(tree)
	assertPaths(t, expected, actual)
}
