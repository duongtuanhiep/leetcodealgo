import math

"""
Question 1201: https://leetcode.com/problems/ugly-number-iii

Since answer is capped at 2 * 10e9, we can possibly generate some number ?
We need to find: 
- LCM of ab,bc,ca and abc
- Each LCMs tell you how many duplicates you have to handle.



"""


class Solution:
    def nthUglyNumber(self, n: int, a: int, b: int, c: int) -> int:

        def uglyCount(n: int) -> int:
            nonlocal a, b, c
            abLCM = math.lcm(a, b)
            cbLCM = math.lcm(b, c)
            caLCM = math.lcm(c, a)
            abcLCM = math.lcm(a, b, c)
            res = (
                n // a
                + n // b
                + n // c
                - n // abLCM
                - n // cbLCM
                - n // caLCM
                + n // abcLCM
            )
            return res

        if min(a, b, c) == 1:
            return n

        lo, hi = 0, 2 * 1e9
        res = 0
        while lo <= hi:
            mid = (hi + lo) // 2
            cur = uglyCount(mid)
            if cur > n:
                hi = mid - 1
            elif cur == n:
                res = mid
                hi = mid - 1
            elif cur < n:
                lo = mid + 1

        return int(res)
