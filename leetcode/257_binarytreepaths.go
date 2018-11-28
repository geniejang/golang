package leetcode

import (
	"strconv"
	"strings"
)

func binaryTreePaths(root *TreeNode) []string {
	stack := []string{}
	return helper(root, stack)
}

func helper(root *TreeNode, stack []string) []string {
	paths := []string{}
	if root != nil {
		height := len(stack)
		stack = append(stack, strconv.Itoa(root.Val))
		paths = append(paths, helper(root.Left, stack)...)
		if root.Left == nil && root.Right == nil {
			paths = append(paths, strings.Join(stack, "->"))
		}
		paths = append(paths, helper(root.Right, stack)...)
		stack = stack[:height]
	}
	return paths
}
