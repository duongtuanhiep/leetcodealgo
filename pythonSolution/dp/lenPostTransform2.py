from typing import List

"""
Question 3337: https://leetcode.com/problems/total-characters-in-string-after-transformations-ii

Constraint:
1 <= s.length <= 1e5
s consists only of lowercase English letters.
1 <= t <= 1e9
nums.length == 26
1 <= nums[i] <= 25

With given contrains, we know that we can represent all character with dp of size 1e9 + 25
"""

""" Attempt with dp + prefix sum: TLE 529/536
class Solution:
    def lengthAfterTransformations(self, s: str, t: int, nums: List[int]) -> int:
        mod = 1e9+7
        dp = [1 for _ in range(26)]
        for i in range(1, t+1):
            prefixSum = [0] * 26
            sum = 0
            for j in range(0, 26):
                sum += dp[j]
                prefixSum[j] = sum

            # print(prefixSum)
            for k in range(0, 26):
                prefSum = 0
                start = k
                end = k + nums[k]
                # print(start, end)
                if end < 26:  # Doesn't wrap
                    prefSum = prefixSum[end] - prefixSum[start]
                else: # Wrap
                    prefSum = prefixSum[25] - prefixSum[start] + prefixSum[end%26]
                prefSum %= mod
                dp[k] = prefSum
            print(dp)

        res = 0 
        for char in s:
            unicode = ord(char) - ord("a")
            res += dp[unicode]
            res %= mod

        return int(res)

"""

""" Base recursion
    def simulate(self, s : str, t: int, nums: List[int]) -> int:
        res = 0 
        for char in s:
            unicode = ord(char) - ord("a")
            if t == 0:
                return nums[unicode]
            for j in range(0, nums[unicode]):
                res += self.simulate(ringAlphabet[unicode+j+1], t-1, nums)
        return res
"""

"""Below is matrix multiplication with exponential optimization"""

MOD = 10**9 + 7
L = 26

class Mat:
    def __init__(self, copy_from: "Mat" = None) -> None:
        self.a: List[List[int]] = [[0] * L for _ in range(L)]
        if copy_from:
            for i in range(L):
                for j in range(L):
                    self.a[i][j] = copy_from.a[i][j]

    def __mul__(self, other: "Mat") -> "Mat":
        result = Mat()
        for i in range(L):
            for j in range(L):
                for k in range(L):
                    result.a[i][j] = (
                        result.a[i][j] + self.a[i][k] * other.a[k][j]
                    ) % MOD
        return result


# identity matrix
def I() -> Mat:
    m = Mat()
    for i in range(L):
        m.a[i][i] = 1
    return m


# matrix exponentiation by squaring
def quickmul(x: Mat, y: int) -> Mat:
    ans = I()
    cur = x
    while y:
        if y & 1:
            ans = ans * cur
        cur = cur * cur
        y >>= 1
    return ans


class Solution:
    def lengthAfterTransformations(
        self, s: str, t: int, nums: List[int]
    ) -> int:
        T = Mat()
        for i in range(26):
            for j in range(1, nums[i] + 1):
                T.a[(i + j) % 26][i] = 1

        res = quickmul(T, t)

        f = [0] * 26
        for ch in s:
            f[ord(ch) - ord("a")] += 1

        ans = 0
        for i in range(26):
            for j in range(26):
                ans = (ans + res.a[i][j] * f[j]) % MOD

        return ans


if __name__ == "__main__":
    cases = [
        ["abcyy", 30, [1,2,1,2,1,2,1,2,1,2,1,2,1,2,1,2,1,2,1,2,1,2,1,2,1,2]], # 7
        # ["azbk", 12, [2,3,2,2,2,2,4,2,2,3,2,2,2,2,2,2,7,2,2,2,2,2,5,2,2,2]], # 16384
    ]

    for c in cases:
        print(Solution().lengthAfterTransformations(c[0],c[1],c[2]))
