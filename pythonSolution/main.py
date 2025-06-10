from typing import Optional
from collections import deque


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def sumEvenGrandparent(self, root: Optional[TreeNode]) -> int:
        res = 0

        # tuple structure (currentNode, parentNode, grandParentNode)
        q = [(root, -1, -1)]
        while q:
            cur, parentVal, grandParentVal = q.pop()
            if cur.left:
                q.append(cur.left, cur.val, parentVal)
            if cur.right:
                q.append(cur.right, cur.val, parentVal)
            if grandParentVal % 2 == 0:
                res += cur.val
        return res
