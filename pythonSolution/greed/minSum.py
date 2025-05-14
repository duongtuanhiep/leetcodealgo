from typing import List

"""

Question 2918:

"""

class Solution:
    def minSum(self, nums1: List[int], nums2: List[int]) -> int:
        s1,s2 = sum(nums1), sum(nums2)
        zero1, zero2 = nums1.count(0), nums2.count(0)
        possibleMinSum = max(s1 + zero1, s2 + zero2)
        if not zero1 and not zero2: 
            return -1 if s1 != s2 else s1
        elif not zero1:
            return -1 if possibleMinSum > s1 else possibleMinSum
        elif not zero2:
            return -1 if possibleMinSum > s2 else possibleMinSum
        return possibleMinSum
