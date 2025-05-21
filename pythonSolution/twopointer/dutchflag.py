from typing import List

"""

Question 75: https://leetcode.com/problems/sort-colors

Dutch flag algorithm.
"""


class Solution:
    def sortColors(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        n = len(nums) - 1
        l, r, i = 0, len(nums) - 1, 0
        while i <= r:
            if nums[i] == 0:
                nums[i], nums[l] = nums[l], nums[i]
                l += 1
                i += 1
            elif nums[i] == 2:
                nums[i], nums[r] = nums[r], nums[i]
                r -= 1
            else:
                i += 1


if __name__ == "__main__":
    cases = [
        [2, 0, 2, 1, 1, 0],
        [2, 0, 1],
        [2, 0, 2, 1, 1, 0, 2, 0, 1, 2, 0, 0, 1, 2, 2, 2],
        [1, 2, 2, 2, 0, 1, 1, 1, 0, 0, 0],
    ]

    for c in cases:
        Solution().sortColors(c)
        print(c)
