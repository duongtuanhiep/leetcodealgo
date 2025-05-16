
"""

Question 3335: https://leetcode.com/problems/total-characters-in-string-after-transformations-i 

Observation:
- each character is independant from others in the input so one can calculate them individually.
- From recursive function, we see a pattern for DP (check un-used function charLen)
- We take that recusive function and flatten it and have a dp solution dp[i] = dp[i-26] + dp[i-25]
- In order to reduce the steps needed when compute dp arrays, we notice that a -> (4steps) == b -> (3 steps) == ...
- This means we only needs to compute z -> (t steps) and that is upperbounded 1e5 + 26

"""


class Solution:
    def lengthAfterTransformations(self, s: str, t: int) -> int:
        # 26 characters * steps 
        total = 26 + t
        mod = 1e9 + 7
        dp = [0] * total
        for i in range(total): 
            if i < 26: 
                dp[i] = 1
            else:
                dp[i] = (dp[i-26] + dp[i-25])%mod

        base = ord("a")
        res = 0
        for char in s:
            pos = ord(char) - base
            res += dp[pos+t]
            res %= mod

        return int(res)

    # Can calculate the length of the final string with input of initial Unicode and steps
    def charLen(self, unicode: int, step: int) -> int:
        res = 0
        if unicode + step < 26:
            return 1

        if unicode+step >= 26:
            res = self.charLen(unicode, step - 26) + self.charLen(unicode+1, step - 26)
        
        return res


if __name__ == "__main__":
    cases = [
        ["abcyy",2], #7
        ["azbk",1], #5
    ]

    for c in cases:
        print(Solution().lengthAfterTransformations(c[0], c[1]))
