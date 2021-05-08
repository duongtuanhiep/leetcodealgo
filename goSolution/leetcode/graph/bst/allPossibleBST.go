package bst

/*
Question 894: https://leetcode.com/problems/all-possible-full-binary-trees/submissions/

Idea:
On each turn at the current node, we can have either 2 child or none. In each 2 child case,
the sub tree can range from 1 to limit (limit is current remaining nodes available to populate).

Also we can do this question iteratively using the idea of "controlled recursion". We can try to fill
each cases by hand (left = 1, right = rest -> left = 3, right = rest -3).

**/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func allPossibleFBT(n int) []*TreeNode {
	if n == 1 {
		return []*TreeNode{&TreeNode{0, nil, nil}}
	} else {
		treeList := []*TreeNode{}
		for i := 1; i < n; i = i + 2 {
			nodeListLeft, nodeListRight := allPossibleFBT(n-i-1), allPossibleFBT(i)
			for _, left := range nodeListLeft {
				for _, right := range nodeListRight {
					newTree := TreeNode{0, left, right}
					treeList = append(treeList, &newTree)
				}
			}
		}
		return treeList
	}
}
