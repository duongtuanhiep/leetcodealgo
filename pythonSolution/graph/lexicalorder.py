from typing import List

"""

Question 386: https://leetcode.com/problems/lexicographical-numbers/

Constraint:
1 <= n <= 5 * 1e4

Ideal: Backtracking DFS


"""


class Solution:
    def lexicalOrder(self, n: int) -> List[int]:
        res = []

        def dfs(start, limit):
            nonlocal res
            for i in range(10):
                if start > 0 or i != 0:  # Appending number
                    next = str(i)
                    if start != 0:
                        next = str(start) + str(i)
                    if int(next) > n:
                        break
                    res.append(int(next))
                    dfs(int(next), limit)

        dfs(0, n)
        return res
