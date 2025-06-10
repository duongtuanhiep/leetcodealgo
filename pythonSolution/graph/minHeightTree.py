from collections import defaultdict
from typing import List

"""

Question 310: https://leetcode.com/problems/minimum-height-trees

Great modification of Kahn algorithm to "trim" all leaf nodes (node with indeg = 1).
We notice that theres can be at most 2 roots that satisfy the requirements.
Kudos to this post: https://leetcode.com/problems/minimum-height-trees/solutions/5060930/full-explanation-bfs-remove-leaf-nodes
"""


class Solution:
    def findMinHeightTrees(self, n: int, edges: List[List[int]]) -> List[int]:
        if n < 3:
            return [i for i in range(n)]

        indegree = [0] * n
        graph = defaultdict(list)
        for e in edges:
            graph[e[0]].append(e[1])
            graph[e[1]].append(e[0])
            indegree[e[0]] += 1
            indegree[e[1]] += 1

        startingLeaf = [i for i in range(n) if indegree[i] == 1]
        processed = 0
        while processed < n - 2:
            cnt = len(startingLeaf)
            # remove current node indegree
            for i in range(cnt):
                leaf = startingLeaf[i]
                indegree[leaf] -= 1
                for next in graph[leaf]:
                    indegree[next] -= 1
                    if indegree[next] == 1:
                        startingLeaf.append(next)
            startingLeaf = startingLeaf[cnt:]
            processed += cnt

        return startingLeaf


if __name__ == "__main__":
    cases = [
        [[[1, 0], [1, 2], [1, 3]], 4],
        [[[3, 0], [3, 1], [3, 2], [3, 4], [5, 4]], 6],
        [[[0, 1], [0, 2]], 3],
    ]

    for c in cases:
        print(Solution().findMinHeightTrees(c[1], c[0]))
