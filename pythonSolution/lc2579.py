"""

Question 2579: https://leetcode.com/problems/count-total-number-of-colored-cells/description/

"""

class Solution:
    def coloredCells(self, n: int) -> int:
        return n**2 * 2 - (n*2-1)
