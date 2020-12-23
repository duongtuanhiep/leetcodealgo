package design

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// Array of pointer with proper in-order traversal
type BSTIterator struct {
	NodeArr []*TreeNode
	Cur     int
}

func Constructor(root *TreeNode) BSTIterator {
	var bst BSTIterator
	bst.NodeArr = traverse(root, bst.NodeArr)
	bst.Cur = 0
	return bst
}

func traverse(node *TreeNode, arr []*TreeNode) []*TreeNode {
	if node.Left != nil {
		arr = traverse(node.Left, arr)
	}
	arr = append(arr, node)
	if node.Right != nil {
		arr = traverse(node.Right, arr)
	}
	return arr
}

func (this *BSTIterator) Next() int {
	this.Cur++
	return this.NodeArr[this.Cur].Val
}

func (this *BSTIterator) HasNext() bool {
	if this.Cur >= len(this.NodeArr)-1 {
		return false
	}
	return true
}

// Stack implementation with controlled recursion
// to "simulate" In order traversal
type BSTIterator struct {
	Stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	var curStack []*TreeNode
	curStack = pushLeft(root, curStack)
	return BSTIterator{Stack: curStack}
}

// push left push all the following left node of current input node, add to stack
func pushLeft(node *TreeNode, arr []*TreeNode) arr {
	for node != nil {
		arr = append(arr, node)
		node = node.Left
	}
	return arr
}

func (this *BSTIterator) Next() int {
	node := this.Stack[len(this.Stack)-1]
	this.Stack = this.Stack[:len(this.Stack)-1]
	this.Stack = pushLeft(node.Right, this.Stack)
	return node
}

func (this *BSTIterator) HasNext() bool {
	return len(this.Stack) != 0
}
