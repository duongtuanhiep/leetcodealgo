from typing import Optional

""" 

Question 143: https://leetcode.com/problems/reorder-list/?envType=problem-list-v2&envId=27ol1gds



"""


# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def reorderList(self, head: Optional[ListNode]) -> None:
        """
        Do not return anything, modify head in-place instead.
        """

        if not head.next:
            return

        old_head = head

        # Count link list node - Possible optimization is to perform i|i*2 iterator to get the division point
        node_count = 1
        while head.next:
            head, node_count = head.next, node_count+1

        # Split link nodes to 2 parts
        prev, head = None, old_head
        for i in range(int(node_count / 2)):
            prev, head = head, head.next
        prev.next = None

        # reverse second part of linked list
        prev, cur, next = None, None, head
        while next:
            if cur:
                cur.next = prev
            next, cur, prev = next.next, next, cur
        cur.next = prev

        # merge 2 parts together
        head, rev_head = old_head, cur
        while head and rev_head:
            tmp1, head.next = head.next, rev_head
            if not tmp1:
                break
            tmp2, rev_head.next = rev_head.next, tmp1
            head, rev_head = tmp1, tmp2

        return old_head

    def print_list(self, node: ListNode | None) -> str:
        res = ""
        while node:
            res += f"{node.val}"
            node = node.next
        return res


if __name__ == "__main__":
    cases = [
        ListNode(1, ListNode(2, ListNode(3, ListNode(4)))),
        ListNode(1, ListNode(2, ListNode(3))),
    ]

    for c in cases:
        Solution().reorderList(c)
