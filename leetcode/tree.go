package leetcode

import (
	"encoding/json"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// newTreeNode returns a TreeNode instance which has val.
// This function is designed to ease description of tests.
func newTreeNode(val string) *TreeNode {
	switch val {
	case "", "null":
		return nil
	}
	num, _ := strconv.Atoi(val)
	return &TreeNode{Val: num, Left: nil, Right: nil}
}

// newTreeFromStringValues returns a root node of built tree from strVals.
// This function is designed to ease description of tests.
func newTreeFromStringValues(strVals string) *TreeNode {
	vals := strings.Split(strVals, ",")
	root := newTreeNode(vals[0])
	lvNodes := []*TreeNode{root}
	for vIdx := 1; vIdx < len(vals); {
		lv := len(lvNodes)
		for nIdx := 0; nIdx < lv; nIdx++ {
			node := lvNodes[nIdx]
			if vIdx < len(vals) {
				node.Left = newTreeNode(vals[vIdx])
				lvNodes = append(lvNodes, node.Left)
			}
			vIdx++
			if vIdx < len(vals) {
				node.Right = newTreeNode(vals[vIdx])
				lvNodes = append(lvNodes, node.Right)
			}
			vIdx++
		}
		lvNodes = lvNodes[lv:]
	}
	return root
}

// JsonString returns json-formatted string of the tree.
// This function is designed to ease description of tests.
func (root *TreeNode) JsonString() string {
	mapTree, _ := json.Marshal(root)
	return string(mapTree)
}
