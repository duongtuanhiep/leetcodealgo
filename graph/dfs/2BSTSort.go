package dfs

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	var arr1 []int
	var arr2 []int
	var res []int
	arr1 = TwoDFS(root1, arr1)
	arr2 = TwoDFS(root2, arr2)
	i, j := 0, 0
	for i < len(arr1) || j < len(arr2) {
		if i == len(arr1) {
			res = append(res, arr2[j])
			j++
		} else if j == len(arr2) {
			res = append(res, arr1[i])
			i++
		} else if arr1[i] < arr2[j] {
			res = append(res, arr1[i])
			i++
		} else {
			res = append(res, arr1[j])
			j++
		}

	}
	return res
}

func TwoDFS(node *TreeNode, arr []int) []int {
	if node == nil {
		return nil
	}

	if node.Left != nil {
		arr = TwoDFS(node.Left, arr)
	}
	arr = append(arr, node.Val)
	if node.Right != nil {
		arr = TwoDFS(node.Right, arr)
	}
	return arr
}
