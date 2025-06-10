from typing import List
from collections import defaultdict, deque

"""

Question 3372: https://leetcode.com/problems/maximize-the-number-of-target-nodes-after-connecting-trees-i/

Precompute the most connected nodes that can be reached with k-1 steps.

"""


class Solution:
    def maxTargetNodes(
        self, edges1: List[List[int]], edges2: List[List[int]], k: int
    ) -> List[int]:

        n, m = len(edges1) + 1, len(edges2) + 1
        graph1, graph2 = defaultdict(list), defaultdict(list)
        for e in edges1:
            graph1[e[0]].append(e[1])
            graph1[e[1]].append(e[0])
        for e in edges2:
            graph2[e[0]].append(e[1])
            graph2[e[1]].append(e[0])

        maxPath1, maxPath2 = [1] * n, [1] * m
        for i in range(n):
            next, visited, steps, score = deque(), set(), 0, 0
            next.append(i)
            while next and steps < k + 1:
                cnt = len(next)
                for _ in range(cnt):
                    cur = next.popleft()
                    if cur in visited:
                        continue
                    visited.add(cur)
                    score += 1
                    for neighbor in graph1[cur]:
                        if neighbor not in visited:
                            next.append(neighbor)
                steps += 1
            maxPath1[i] = score

        for i in range(m):
            next, visited, steps, score = deque(), set(), 0, 0
            next.append(i)
            while next and steps < k:
                cnt = len(next)
                for _ in range(cnt):
                    cur = next.popleft()
                    if cur in visited:
                        continue
                    visited.add(cur)
                    score += 1
                    for neighbor in graph2[cur]:
                        if neighbor not in visited:
                            next.append(neighbor)
                steps += 1
            maxPath2[i] = score

        bestNode = max(maxPath2)
        return [maxPath1[i] + bestNode for i in range(n)]
