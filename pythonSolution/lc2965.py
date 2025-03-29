
"""

Question 2965: https://leetcode.com/problems/find-missing-and-repeated-values/

"""


class Solution:
    def findMissingAndRepeatedValues(self, grid: List[List[int]]) -> List[int]:
        count = [0 * i for i in range(len(grid) ** 2) ]
        for i in range(len(grid)):
            for j in range(len(grid[i])):
                count[grid[i][j]-1] += 1

        dup, miss = -1, -1
        for i, val in enumerate(count):
            if val > 1: 
                dup = i+1
            if val == 0:
                miss = i+1

        return [dup,miss]
