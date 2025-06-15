from typing import List


"""

Question 1011: 

"""


class Solution:

    # Binary search
    def shipWithinDays(self, weights: List[int], days: int) -> int:
        hi = 500 * 1e4 * 5
        lo = max(weights)

        def canShip(shipCapacity, days):
            nonlocal weights
            cur = 0
            for w in weights:
                if cur + w <= shipCapacity:
                    cur += w
                else:
                    cur = w
                    days -= 1
            return days >= 1

        res = hi
        while lo <= hi:
            mid = lo + (hi - lo) // 2
            if canShip(mid, days):
                res = mid
                hi = mid - 1
            else:
                lo = mid + 1

        # This works too
        # while lo < hi:
        #     mid = lo + (hi - lo) // 2
        #     if canShip(mid, days):
        #         res = mid
        #         hi = mid
        #     else:
        #         lo = mid + 1

        return int(res)
