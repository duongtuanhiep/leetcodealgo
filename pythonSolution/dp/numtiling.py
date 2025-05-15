import functools

"""
Question 790: https://leetcode.com/problems/domino-and-tromino-tiling


This require carefully chosen base cases. Best explanation I found is here:
https://leetcode.com/problems/domino-and-tromino-tiling/solutions/6715436/recursion-tabulation-math-bonus-o-1-at-end-with-images-hashmap-c-python-java
"""


class Solution(object):
    def numTilings(self, n):
        mod = 10**9 + 7
        if n <= 1:
            return 1
        if n == 2:
            return 2
        if n == 3:
            return 5
        dp = [0] * (n + 1)
        dp[0], dp[1], dp[2], dp[3] = 1, 1, 2, 5
        for i in range(4, n+1):
            dp[i] = (dp[i-1]*2 + dp[i-3]) % mod
        return dp[n]
