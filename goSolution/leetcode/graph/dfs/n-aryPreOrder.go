package dfs

/*
Question 589: https://leetcode.com/problems/n-ary-tree-preorder-traversal/

Solution:
- Implement recursion
- Implement stack
*/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}

	res := []int{root.Val}
	for _, node := range root.Children {
		res = append(res, preorder(node)...)
	}
	return res
}

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	stack, res := []*Node{}, []int{}
	stack = append(stack, root)
	for len(stack) > 0 {
		res = append(res, stack[0].Val)
		newStack := []*Node{}
		for _, childNode := range stack[0].Children {
			newStack = append(newStack, childNode)
		}
		stack = stack[1:]
		stack = append(newStack, stack...)
	}

	return res
}
