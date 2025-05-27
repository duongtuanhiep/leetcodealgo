from typing import List
from collections import defaultdict

"""

Question 1857: https://leetcode.com/problems/largest-color-value-in-a-directed-graph

Tricks to check for cycle detection: each node has 3 states 0,1,2
Instead of worriyng which child to pick, we pick "all" with max of each possible color
"""


class Solution:
    def largestPathValue(self, colors: str, edges: List[List[int]]) -> int:
        visited = [0] * len(colors)
        colorCnt = [[0 for _ in range(26)] for _ in range(len(colors))]
        res = -1

        graph = defaultdict(list)
        for e in edges:
            graph[e[0]].append(e[1])

        def dfs(node: int) -> int:
            color = ord(colors[node]) - ord("a")

            if visited[node] == 1:
                return -1
            elif visited[node] == 2:
                return colorCnt[node][color]
            visited[node] = 1

            for next in graph[node]:
                if visited[next] == 1:
                    return -1

                dfs(next)
                for i in range(26):
                    colorCnt[node][i] = max(colorCnt[node][i], colorCnt[next][i])

            colorCnt[node][color] += 1
            visited[node] = 2
            return colorCnt[node][color]

        for i in range(len(colors)):
            cur = dfs(i)
            if cur > res:
                res = cur
            if cur == -1:
                return -1

        return res
