
# Question 326
# https://leetcode.com/problems/power-of-three/
# First question done in python 3 : - ) 

class Solution:
    def isPowerOfThree(self, n: int) -> bool:
        if n == 1:
            return True
        elif n % 3 == 0 and n >= 3:
            return Solution.isPowerOfThree(self,n/3)
        return False