"""
Question 3443: https://leetcode.com/problems/maximum-manhattan-distance-after-k-changes/
"""


class Solution:
    def maxDistance(self, s: str, k: int) -> int:
        """
        Manhattan distance: abs(x) + abs(y)
        Swap K so that max(n,s) + max(w,e) | It's also independent between N/S and W/E
        """
        north, west, east, south = 0, 0, 0, 0
        res = 0
        for d in s:
            if d == "N":
                north += 1
            elif d == "S":
                south += 1
            elif d == "W":
                west += 1
            elif d == "E":
                east += 1

            # dist without k changes
            currentBest = abs(north - south) + abs(west - east)

            # dist with k changes
            currentBest += min(min(north, south) + min(west, east), k) * 2

            res = max(res, currentBest)

        return res
