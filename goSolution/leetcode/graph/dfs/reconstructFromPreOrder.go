package dfs

import (
	"fmt"
	"strconv"
	"strings"
)

/*

Question 1028: https://leetcode.com/problems/recover-a-tree-from-preorder-traversal/

Since the order of the traversal is pre order: it means for a given string representation,
there will be these component:
- [depth]CurrentValues | [depth-1]Left | [depth-1]Right

We can devise a recursive approach where on each iteration, we:
- keep the current "depth" to calculate the number of leading '-'
- Try to split the current string into 3 parts: current node, possible left subtree and possible right subtree
- (optimization) instead of passing the whole subtree into recursion stack we pass only start, end index in the original arr

**/

func recoverFromPreorder(traversal string) *TreeNode {
	// pre-build the traversal string
	// count 'number of continuous hyphens in each appearance'

	for i := 0; i < len(traversal); i++ {
		if traversal[i] == '-' {
			var hyphenCount int
			for j := i; j < len(traversal); j++ {
				if traversal[j] == '-' {
					hyphenCount++
				} else {
					break
				}
			}
			// modify traversal, replacing hyphen with a count
			traversal = traversal[:i] + fmt.Sprintf("_%v!", hyphenCount) + traversal[i+hyphenCount:]
		}
	}

	traversal = "_0!" + traversal // adding 0 depth
	return recover(traversal, 0)
}

func recover(input string, depth int) *TreeNode {
	if len(input) <= 0 {
		return nil
	}

	nextDepthPrefix := fmt.Sprintf("_%v!", depth+1)
	currentDepthPrefix := fmt.Sprintf("_%v!", depth)
	childIndex := strings.Index(input, nextDepthPrefix)
	if childIndex == -1 { // only has current node
		curNum, _ := strconv.Atoi(input[len(currentDepthPrefix):])
		return &TreeNode{Val: curNum}
	} else {
		var leftNode, rightNode *TreeNode

		curNum, _ := strconv.Atoi(input[len(currentDepthPrefix):childIndex])

		parts := strings.Split(input, nextDepthPrefix)
		if len(parts) > 2 {
			leftNode = recover(nextDepthPrefix+parts[1], depth+1)
			rightNode = recover(nextDepthPrefix+parts[2], depth+1)
		} else {
			leftNode = recover(nextDepthPrefix+parts[1], depth+1)
		}
		return &TreeNode{Val: curNum, Left: leftNode, Right: rightNode}
	}
}
