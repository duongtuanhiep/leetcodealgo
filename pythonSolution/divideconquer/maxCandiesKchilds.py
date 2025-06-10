from typing import List

"""

Question 2226: https://leetcode.com/problems/maximum-candies-allocated-to-k-children/

Solutions: 
- Presort input arrays
- Perform BS with hi = max(arrays[n]) - maximum candy available, lo = 0
- For each input, if can_distribute(), we try to search up. 

"""

class Solution:
    def maximumCandies(self, candies: List[int], k: int) -> int:
        candies.sort(reverse=True)
        res, hi, lo = 0, candies[0], 1
        while lo <= hi:
            target = (hi + lo) // 2
            if self.can_distribute(candies, target, k):
                res = target
                lo = target + 1
            else:
                hi = target - 1

        return res

    def can_distribute(self, candies: List[int], target, k: int) -> bool:
        for candy_count in candies:
            if candy_count < target or k <= 0:
                break
            k -= candy_count // target

        return k <= 0

if __name__ == "__main__":
    cases = [
        # [[5,8,6], 3],
        [[6,6,6,6,6], 11],
    ]

    for case in cases:
        print(Solution().maximumCandies(case[0], case[1]))
