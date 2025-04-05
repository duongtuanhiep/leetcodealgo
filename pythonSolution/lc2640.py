"""

Question 2640: https://leetcode.com/problems/apply-operations-to-an-array/

"""


class Solution:
    def applyOperations(self, nums: List[int]) -> List[int]:
        for i in range(len(nums)):
            if i + 1 < len(nums) and nums[i] == nums[i + 1]:
                nums[i] *= 2
                nums[i + 1] = 0

        result, j = [0] * len(nums), 0
        for i in range(len(nums)):
            if nums[i] != 0:
                result[j], j = nums[i], j + 1
        return result
