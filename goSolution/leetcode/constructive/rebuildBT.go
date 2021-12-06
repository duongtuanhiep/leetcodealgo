package constructive

/**
Question 105:
https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }

Test:
[3,9,20,15,7,86,95]
[9,3,15,20,7,95,86]
[1]
[1]
[1,2]
[2,1]
[1,2]
[1,2]
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	locIn, locPre := make(map[int]int), make(map[int]int)
	for i, val := range inorder {
		locIn[val] = i
	}
	for i, val := range preorder {
		locPre[val] = i
	}

	return build(preorder, inorder, len(preorder)-1, 0, locIn, locPre)
}

func build(pre, in []int, hi, lo int, locIn, locPre map[int]int) *TreeNode {
	if lo > hi {
		return nil
	} else if lo == hi {
		return &TreeNode{pre[lo], nil, nil}
	}

	curNode := &TreeNode{pre[lo], nil, nil}
	var leftLo, leftHi, rightLo, rightHi int
	leftLo = lo + 1
	if locIn[curNode.Val]-1 < 0 {
		leftHi = -1
	} else {
		leftHi = locPre[in[locIn[curNode.Val]-1]]
	}
	rightLo = locIn[curNode.Val] + 1
	rightHi = hi
	curNode.Left = build(pre, in, leftHi, leftLo, locIn, locPre)
	curNode.Right = build(pre, in, rightHi, rightLo, locIn, locPre)
	return curNode
}
