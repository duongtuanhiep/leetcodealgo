from typing import Optional

"""

Question 138: https://leetcode.com/problems/copy-list-with-random-pointer

Idea:
- Very big brain solutions: Add duplicates elements right after real elements, move the ref manually and then link the copied nodes to the copieds one.
- Or you can do it recursively like a graph problems, just need to have a set to check which one you've visited.
"""


# Definition for a Node.
class Node:
    def __init__(self, x: int, next: "Node" = None, random: "Node" = None):
        self.val = int(x)
        self.next = next
        self.random = random


class Solution:
    def copyRandomList(self, head: "Optional[Node]") -> "Optional[Node]":
        dummy = Node(-1, head, None)

        if not head:
            return None

        # Create dup of each node in the original list
        # A -> A' -> B -> B' -> C -> C' -> NULL
        while head:
            # Duplicate nodes points to the original next/random nodes
            clone = Node(head.val, head.next, head.random)

            # Move original node's next to the new clone nodes - we don't lose anything since clone also points to og next node
            clone.next, head.next = head.next, clone

            # Move past the clone node to the original's next
            head = head.next.next

        # New list's head
        head = dummy.next.next

        # Move next and random pointers to next
        # Since previous dups still had reference to original node, now we move it to the copy node.
        while head:
            # Move next from original to copy node
            if head.next:
                head.next = head.next.next

            # Move random from original to copy node
            if head.random:
                head.random = head.random.next

            head = head.next

        return dummy.next.next
