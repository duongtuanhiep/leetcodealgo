

"""
Goodpair only happens if nums[i] == nums[j] 

"""

class Solution:
    def numIdenticalPairs(self, nums: List[int]) -> int:
        bucket = [0 * i for i in range(100)]
        res = 0
        for val in nums:
            bucket[val-1] += 1
            
        for counts in bucket:
            if counts >= 2:
                res += counts * (counts-1) / 2

        return int(res)




