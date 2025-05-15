from typing import Optional

"""

Question 124: https://leetcode.com/problems/binary-tree-maximum-path-sum

"""

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:

    def maxPathSum(self, root: Optional[TreeNode]) -> int:
        res = -2000
        def dfs(node: TreeNode | None ) -> int: # Return the best path through that node
            nonlocal res  # Declare that 'res' is from the outer scope

            if node == None:
                return -2000
            
            curVal = node.val
            left, right = dfs(node.left), dfs(node.right)
            curMax = curVal + max(left, 0) + max(right, 0) # best path "turn" at current node
            subMax = curVal + max(left, right, 0) # best path "turn" at some parent node above
            res = curMax if curMax > res else res
            return subMax

        dfs(root)
        return res
