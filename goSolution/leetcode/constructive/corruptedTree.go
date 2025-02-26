package constructive

/*

Question 1261: https://leetcode.com/problems/find-elements-in-a-contaminated-binary-tree/description/

**/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type FindElements struct {
	exist map[int]struct{}
}

func Constructor(root *TreeNode) FindElements {
	fe := FindElements{exist: make(map[int]struct{})}
	root.Val = 0
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		cur := queue[0]
		fe.exist[cur.Val] = struct{}{}

		if cur.Left != nil {
			cur.Left.Val = 2*cur.Val + 1
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			cur.Right.Val = 2*cur.Val + 2
			queue = append(queue, cur.Right)
		}

		queue = queue[1:]
	}

	return fe
}

func (this *FindElements) Find(target int) bool {
	_, found := this.exist[target]
	return found
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
