"""

Question 2206: https://leetcode.com/problems/divide-array-into-equal-pairs

You could also use counter=collections.Counter(nums)

"""

class Solution:
    def divideArray(self, nums: List[int]) -> bool:
        count = [0 for _ in range(501)]
        for n in nums:
            count[n] += 1 

        for c in count:
            if c % 2 != 0:
                return False
        return True
