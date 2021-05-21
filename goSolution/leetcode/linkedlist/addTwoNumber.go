package linkedlist

/*
Question 2: https://leetcode.com/problems/add-two-numbers/

Idea:
 - Store a the possible node on an array, returns the index i = 1 of
 that array
**/
/*
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var holder int
	resArr := []*ListNode{&ListNode{Val: 10, Next: nil}}
	for l1 != nil || l2 != nil || holder != 0 {
		var val1, val2 int
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}
		var curNode ListNode
		val := val1 + val2 + holder
		curNode.Val = val % 10
		holder = (val - val%10) / 10
		resArr[len(resArr)-1].Next = &curNode
		resArr = append(resArr, &curNode)
	}
	if len(resArr) > 1 {
		return resArr[1]
	}
	return nil
}
*/
/*
- Perharps we can create a symbolic initial node since it is guaranteed
that it will not be zero and below 10, can assign an abitrary value outside
acceptable scope of Val and initialize based of that "fake Node"
**/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var holder int
	dummyNode := ListNode{Val: -1, Next: nil}
	symbolNode := &dummyNode
	for l1 != nil || l2 != nil || holder != 0 {
		var val1, val2 int
		if l1 != nil {
			val1, l1 = l1.Val, l1.Next
		}
		if l2 != nil {
			val2, l2 = l2.Val, l2.Next
		}
		sum := val1 + val2 + holder
		var curNode ListNode
		curNode.Val = sum % 10
		holder = (sum - curNode.Val%10) / 10
		symbolNode.Next = &curNode
		symbolNode = symbolNode.Next
	}
	if dummyNode.Next != nil {
		return dummyNode.Next
	}
	return nil
}
