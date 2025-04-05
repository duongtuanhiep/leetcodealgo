# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

"""
Some detection rules:
- if next to last layer, should only has left

"""


class Solution:
    def isCompleteTree(self, root: Optional[TreeNode]) -> bool:
        queue = [root]
        isLastLayer = False
        while queue:
            final = isLastLayer
            currLen = len(queue)
            for i in range(currLen):  # Should have 2 or left
                curr, queue = queue[0], queue[1:]

                if isLastLayer and (curr.left or curr.right):
                    return False

                hasLeft = False
                if curr.left:
                    queue.append(curr.left)
                    hasLeft = True
                else:
                    isLastLayer = True

                if curr.right:
                    if hasLeft:
                        queue.append(curr.right)
                    else:
                        return False
                else:
                    isLastLayer = True

            if final and len(queue) > 0:
                return False

        return True
