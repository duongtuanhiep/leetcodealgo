package bfs

/*
Question 968: https://leetcode.com/problems/binary-tree-cameras/


Test:
[0,0,null,0,null,0,null,null,0]
[0,0,null,0,0]

Solution: https://leetcode.com/problems/binary-tree-cameras/solution/
Idea :
Greedy:
 - It is strictly better to put cam on parent if child node does not have cam and not covered.
 - We create a map to keep track of which node has been covered.
 - From the initial observation we notice one way to do it is to check for last parent (suggest DFS, go to deepest layer first)
 - Keep track of the last since if we put the cam at current node the last should also be covered.

Dynamic:

**/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minCameraCover(root *TreeNode) int {
	covered := make(map[*TreeNode]bool)
	return cover(root, nil, covered)
}

func cover(cur, last *TreeNode, covered map[*TreeNode]bool) int {
	var res int
	if cur.Left != nil {
		res += cover(cur.Left, cur, covered)
	}
	if cur.Right != nil {
		res += cover(cur.Right, cur, covered)
	}

	if !covered[cur.Left] && cur.Left != nil || !covered[cur.Right] && cur.Right != nil || last == nil && !covered[cur] {
		res++
		covered[cur] = true
		if cur.Left != nil {
			covered[cur.Left] = true
		}
		if cur.Right != nil {
			covered[cur.Right] = true
		}
		if last != nil {
			covered[last] = true
		}
	}

	return res
}
