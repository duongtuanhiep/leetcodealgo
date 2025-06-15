from typing import List
import heapq

"""

Question 135: https://leetcode.com/problems/candy

O(NlogN) time + O(N) space
- Greedy with heapQ solutions: Pick the lowest rating child first and assign the least possible amount of candies for him. 

O(N): DP/Greedy
- Could possibly solve this with dp too where dp[i] = dp[i-1] + 1 | dp[i+1] + 1

"""


class Solution:
    def candy1(self, ratings: List[int]) -> int:
        hq = []
        for index, rating in enumerate(ratings):
            heapq.heappush(hq, (rating, index))

        res = [1] * len(ratings)
        while hq:
            candies = 1
            lowestRating, lowestIndex = heapq.heappop(hq)
            if lowestIndex - 1 >= 0 and lowestRating > ratings[lowestIndex - 1]:
                candies = res[lowestIndex - 1] + 1
            if (
                lowestIndex + 1 < len(ratings)
                and lowestRating > ratings[lowestIndex + 1]
            ):
                candies = max(candies, res[lowestIndex + 1] + 1)
            res[lowestIndex] = candies
        return sum(res)

    def candy(self, ratings):
        n = len(ratings)

        dp = [1] * n

        for i in range(1, n):
            if ratings[i] > ratings[i - 1]:
                dp[i] = max(dp[i], 1 + dp[i - 1])

        for i in range(n - 2, -1, -1):
            if ratings[i] > ratings[i + 1]:
                dp[i] = max(dp[i], 1 + dp[i + 1])

        return sum(dp)
