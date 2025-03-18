import math

"""

Question 1780: https://leetcode.com/problems/check-if-number-is-a-sum-of-powers-of-three/


// Magic arrays [1, 3, 9, 27, 81, 243, 729, 2187, 6561, 19683, 59049, 177147, 531441, 1594323]

"""


class Solution:
    def checkPowersOfThree(self, n: int) -> bool:
        pows = []
        for i in range(0,n+1,1):
            if 3**i > n: 
                break
            pows.append(3**i)

        used = set()
        while n > 0:
            sumCandidate = self.binSearch(n, pows)
            if sumCandidate in used: 
                return False
            else:
                n -= sumCandidate
                used.add(sumCandidate)

        return n == 0


    # get the greatest smaller number to n 
    def binSearch(self, n: int, arr: List[int]) -> int:
        res = -1
        hi,lo = len(arr)-1, 0
        while lo <= hi:
            mid = (hi+lo)//2
            if arr[mid] <= n:
                res = arr[mid]
                lo = mid +1
            else: 
                hi = mid -1
        return  res
