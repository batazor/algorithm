package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil && q != nil || p != nil && q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	left := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
			},
		},
	}

	right := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
			},
		},
	}

	fmt.Println(isSameTree(left, right)) // false
	fmt.Println(isSameTree(left, left))  // true
}
