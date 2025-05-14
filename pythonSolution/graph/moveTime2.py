from typing import List
import heapq

"""

Question 3342: https://leetcode.com/problems/find-minimum-time-to-reach-last-room-ii

"""
directions = [(1, 0), (-1, 0), (0, 1), (0, -1)]


class Solution:
    def minTimeToReach(self, moveTime: List[List[int]]) -> int:
        m, n = len(moveTime), len(moveTime[0])
        visitCost = [[-1 for _ in range(n)] for _ in range(m)]
        pq = []
        heapq.heappush(pq, (0, 2, 0, 0))
        while len(pq) > 0:
            cost, travel, i, j = heapq.heappop(pq)
            if visitCost[i][j] != -1:
                continue
            visitCost[i][j] = cost
            for x, y in directions:
                nextI, nextJ = i + x, j + y
                if 0 <= nextI < m and 0 <= nextJ < n and visitCost[nextI][nextJ] == -1:
                    nextTravel = 2 if travel == 1 else 1
                    nextCost = max(
                        cost + nextTravel, moveTime[nextI][nextJ] + nextTravel
                    )
                    heapq.heappush(pq, (nextCost, nextTravel, nextI, nextJ))

        return visitCost[m - 1][n - 1]
