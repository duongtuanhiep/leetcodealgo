from typing import List
from collections import deque

"""

Question 1425: https://leetcode.com/problems/constrained-subsequence-sum

Solution: Dynamic Programming O(NlogN)
- Keeps a priority queue of the last sub arrays ends at previous k->i steps, on each steps, it's dp[i] = max(heap.Max(dp[i-k:i]) + nums[i], nums[i])

Solutions: Strictly decreasing monotonic queue.
- Instead of keeping a heap and takes LogN to gets the best possible candidates, we use a strictly decreasing monotonic queue.
- If we can make sure that the queue is strictly decreasing, this means that best possible sum that current index i can get will be at queue[0].
- Strictly decreasing because: If we have queue=[a,b,c] and new items d is greater than b and c, theres no need to keep b,c for next iteration since d is better. Queue would then be [a,d]
- On everysteps, result would be current = max(nums[i], queue[0] + nums[i]) if queue is not empty.
"""


class Solution:
    def constrainedSubsetSum(self, nums: List[int], k: int) -> int:
        # Monotonic queue and a sum.
        res = -1e4 - 1
        mq = deque()  # Storing (bestCurrentSum, x) with x in range [i-k:i]
        rollingSum = 0

        for i, n in enumerate(nums):
            # Cleaning up items that are not accessible in the current iteration
            while mq and mq[0][1] + k < i:
                mq.popleft()

            bestCurrent = n if not mq else max(n, n + mq[0][0])
            res = max(res, bestCurrent)

            while mq and bestCurrent > mq[-1][0]:
                mq.pop()
            mq.append((bestCurrent, i))

        return res
