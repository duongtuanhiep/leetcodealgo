from typing import List
from collections import deque

"""

Question 42: https://leetcode.com/problems/trapping-rain-water

Very typical usecase of a monotonic stack (decreasing in this case).
We greadily tries to add more element to the queue, if an element more than last element then we have potential "rain water". 

"""


class Solution:
    def trap(self, height: List[int]) -> int:
        res = 0
        stack = deque()
        lastWaterHeight, lastWaterIndex = -1, -1
        for i, h in enumerate(height):
            # Pop when not decreasing
            while len(stack) > 0 and stack[-1][0] < h:
                stack.pop()

                # Calculate extra rain water
                if len(stack) > 0:
                    start, end = stack[-1][1], i
                    waterHeight = min(h, stack[-1][0])

                    # recalculate based on prev calculation
                    if start <= lastWaterIndex and waterHeight > lastWaterHeight:
                        res += (end - start - 1) * (waterHeight - lastWaterHeight)
                    # Fresh calculation if start dooesn't overlap with previous water index
                    elif start > lastWaterIndex:
                        for j in range(start, end):
                            res += max(waterHeight - height[j], 0)

                    lastWaterIndex, lastWaterHeight = start, waterHeight

            stack.append((h, i))

        return res


if __name__ == "__main__":
    cases = [
        [0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1],  # 6
        [4, 2, 0, 3, 2, 5],  # 9
        [0, 0, 0, 4, 3, 2, 1, 2, 3, 4, 5, 0, 0, 1],  # 11
    ]

    for c in cases:
        print(Solution().trap(c))
