package leetcode

func isValidBST(root *TreeNode) bool {
	node := root
	stack, size := []*TreeNode{}, 0
	prev, isFirst := 0, true
	for len(stack) > 0 || node != nil {
		if node != nil {
			stack, size = append(stack, node), size+1
			node = node.Left
		} else {
			size--
			node, stack = stack[size], stack[:size]
			if isFirst {
				isFirst = false
			} else if node.Val <= prev {
				return false
			}
			prev = node.Val
			node = node.Right
		}
	}
	return true
}
