package dynamicp

/*
Question 124: https://leetcode.com/problems/binary-tree-maximum-path-sum/

Idea:
1) If we can get max path that has root at each node we can get max path of whole tree.
2) Each max path at current node can be either:
- Only includes itself
- Include left
- Include right
- Include both (if include both it wont be any anynode childs)
3) store the solution at every spot as a pair of int. First only contains left or right along with itself, second contains both child

Good explanation:
https://leetcode.com/problems/binary-tree-maximum-path-sum/discuss/603423/Python-Recursion-stack-thinking-process-diagram
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
Testcase:
[5,4,8,11,null,13,4,7,2,null,null,null,1]
[-10,9,20,null,null,15,7]
*/

func maxPathSum(root *TreeNode) int {
	memo := make(map[*TreeNode][]int)
	maxSumInclSelf(root, memo)
	res := -1001
	for _, valPair := range memo {
		if valPair[1] > res {
			res = valPair[1]
		}
	}
	return res
}

func maxSumInclSelf(root *TreeNode, memo map[*TreeNode][]int) {
	if root.Left != nil {
		maxSumInclSelf(root.Left, memo)
	}
	if root.Right != nil {
		maxSumInclSelf(root.Right, memo)
	}
	valPair := []int{root.Val, root.Val}

	// Cal root
	var left, right int
	if root.Left != nil {
		left = memo[root.Left][0]
	}
	if root.Right != nil {
		right = memo[root.Right][0]
	}

	if max(left, right) > 0 {
		valPair[0] += max(left, right)
	}
	if left > 0 {
		valPair[1] += left
	}
	if right > 0 {
		valPair[1] += right
	}

	memo[root] = valPair
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
