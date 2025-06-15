"""

Question 236: https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree

Idea is that we should return the first node where either itself, left, right trees contain both p and q
Naturally this leads to DFS

A more elegants ways to do this is to have mid,left,right

"""


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def lowestCommonAncestor(
        self, root: "TreeNode", p: "TreeNode", q: "TreeNode"
    ) -> "TreeNode":
        dummy = root

        def dfs(
            node: TreeNode, p: TreeNode, q: TreeNode
        ) -> tuple[TreeNode, bool, bool]:
            hasP = True if node == p else False
            hasQ = True if node == q else False

            if node.left:
                candidateLeft, hasPLeft, hasQLeft = dfs(node.left, p, q)
                if hasPLeft and hasQLeft:
                    return candidateLeft, True, True

            if node.right:
                candidateRight, hasPRight, hasQRight = dfs(node.right, p, q)
                if hasPRight and hasQRight:
                    return candidateRight, True, True

            hasP = hasP or hasPLeft or hasPRight
            hasQ = hasQ or hasQLeft or hasQRight
            if hasQ and hasP:
                return node, True, True

            return None, hasQ, hasP

        res, _, _ = dfs(root, p, q)
        return res
