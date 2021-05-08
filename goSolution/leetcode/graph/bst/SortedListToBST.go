package bst

/**
Question 109:
https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	return makeTree(arr, len(arr)-1, 0)
}

func makeTree(arr []int, hi, lo int) *TreeNode {
	if lo > hi {
		return nil
	}
	avg := (hi + lo) / 2
	root := &TreeNode{arr[avg], nil, nil}
	root.Left, root.Right = makeTree(arr, avg-1, lo), makeTree(arr, hi, avg+1)
	return root
}
