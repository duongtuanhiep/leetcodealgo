from typing import List

"""

Question 55:https://leetcode.com/problems/jump-game

"""


class Solution:

    # Generic DFS
    def canJumpDFS(self, nums: List[int]) -> bool:
        visited = set()
        nextIndex = [0]
        while nextIndex:
            cur = nextIndex.pop()
            if cur in visited:
                continue
            visited.add(cur)

            for i in range(1, nums[cur] + 1):
                next = i + cur
                if next >= len(nums):
                    break
                if next not in visited:
                    nextIndex.append(i + cur)

        return len(nums) - 1 in visited

    # Greedy approach with 1 checks
    def canJump(self, nums: List[int]) -> bool:
        lastGoodIndex = len(nums) - 1

        for i in range(lastGoodIndex - 1, -1, -1):
            if i + nums[i] >= lastGoodIndex:
                lastGoodIndex = i

        return lastGoodIndex == 0
