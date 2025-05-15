from typing import List
import heapq

"""

Question 3341: https://leetcode.com/problems/find-minimum-time-to-reach-last-room-i

Implement djikstra
"""

directions = [(1, 0), (-1, 0), (0, 1), (0, -1)]

class Solution:
    def minTimeToReach(self, moveTime: List[List[int]]) -> int:
        m, n = len(moveTime), len(moveTime[0])
        visited = [[-1 for _ in range(n)] for _ in range(m)]
        pq = []
        heapq.heappush(pq, (0, 0, 0))

        while len(pq) > 0:
            cost, i, j = heapq.heappop(pq)
            if visited[i][j] != -1:
                continue
            visited[i][j] = cost
            for a, b in directions:
                newCost, newI, newJ = cost + 1, i + a, j + b
                if 0 <= newI < m and 0 <= newJ < n and visited[newI][newJ] == -1:
                    newCost = max(newCost, moveTime[newI][newJ] + 1)
                    heapq.heappush(pq, (newCost, newI, newJ))

        return visited[m - 1][n - 1]


if __name__ == "__main__":
    cases = [
        [[0, 4], [4, 4]],  # 6
        [[0, 0, 0], [0, 0, 0]],  # 3
        [[0, 0, 0], [0, 9, 9], [0, 0, 0]],  # 4
        [[0, 1], [1, 2]],  # 3
    ]

    for c in cases:
        print(Solution().minTimeToReach(c))
