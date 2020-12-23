package graph

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//
func isSubtree(s *TreeNode, t *TreeNode) bool {
	var nodeLists []*TreeNode
	sortedStack := traverseInDFS(s, nodeLists)
	answer := false
	copyOfSub := t
	for pos := range sortedStack {
		if coolDoubleDFS(sortedStack[pos], copyOfSub) {
			answer = true
			break
		}
	}
	return answer
}

func traverseInDFS(node *TreeNode, stacks []*TreeNode) []*TreeNode {
	if node.Left != nil {
		stacks = traverseInDFS(node.Left, stacks)
	}
	stacks = append(stacks, node)
	if node.Right != nil {
		stacks = traverseInDFS(node.Right, stacks)
	}
	return stacks
}

func coolDoubleDFS(main *TreeNode, sub *TreeNode) bool {
	answer1, answer2 := true, true
	if main.Left != nil && sub.Left != nil {
		answer1 = coolDoubleDFS(main.Left, sub.Left)
	} else if main.Left == nil && sub.Left != nil || main.Left != nil && sub.Left == nil {
		answer1 = false
	}
	if main.Val != sub.Val {
		answer1 = false
	}
	if main.Right != nil && sub.Right != nil {
		answer2 = coolDoubleDFS(main.Right, sub.Right)
	} else if main.Right == nil && sub.Right != nil || main.Right != nil && sub.Right == nil {
		answer2 = false
	}
	if answer1 == answer2 && answer1 == true {
		return true
	}
	return false
}

// func traverseCompare(first *TreeNode, second *TreeNode) bool {}
