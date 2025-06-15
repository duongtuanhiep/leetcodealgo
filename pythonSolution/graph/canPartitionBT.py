from typing import Optional

"""

Question 633: https://leetcode.com/problems/equal-tree-partition

Observations: 
- Determine if subTrees from current node has half the sum values

"""


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def checkEqualTree(self, root: Optional[TreeNode]) -> bool:

        # DFS with a stack to get max score
        totalScore = 0
        q = [root]
        while q:
            cur = q.pop()
            totalScore += cur.val
            if cur.left:
                q.append(cur.left)
            if cur.right:
                q.append(cur.right)

        canPartition = False

        def dfsSum(node) -> int:
            nonlocal root, canPartition, totalScore

            nodeSum = node.val
            if node.left:
                nodeSum += dfsSum(node.left)
            if node.right:
                nodeSum += dfsSum(node.right)
            if nodeSum * 2 == totalScore and node != root:
                canPartition = True

            return nodeSum

        dfsSum(root)
        return canPartition
