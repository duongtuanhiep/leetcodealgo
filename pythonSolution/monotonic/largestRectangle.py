from typing import List

"""

Question 84: https://leetcode.com/problems/largest-rectangle-in-histogram

Key observation: 
- For each height at index x, we want to calculate the largest rectangle with height[x]. 
- For this, we want to find the maximum width, ie the leftmost bound and the rightmost bound of the rectangle with height[x]. 

This is when the monotonically increasing stack comes in:
- We maintain monotonically increasing property. 
- For each index i:
    - if heights[i] greater that last in the stack, we add it to the stack. 
    - if heights[i] smaller than the last element in the stack, this will be the rightmost bound of ALL the element x with height[x] greater than height[i] in our stack
        - We then pop each of the elements x with heights[x] > heights[i]:
        - The leftmost bound of the popped elements will be then-last-element in the stack (because it is monotonically increasing !). 
        - We calculate the area and then compare it with res. 

Smart optimization: 
- In order to not having to deal with left-over elements from the end & beginning, we artificially add our own bounds with -1 height to both end.

"""


class Solution:
    def largestRectangleArea(self, heights: List[int]) -> int:
        res = 0
        ms = []  # Item: (height, index)
        for i, n in enumerate([-1] + heights + [-1]):
            while ms and ms[-1][0] > n:
                lastHeight, _ = ms.pop()
                res = max(res, lastHeight * (i - ms[-1][1] - 1))
            ms.append((n, i))

        return res
