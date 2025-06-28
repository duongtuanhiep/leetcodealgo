from typing import List


"""

Question 2551: https://leetcode.com/problems/put-marbles-in-bags

One singular observation changes everythings: result always contains w[0] + w[1] and k-1 "cutting point"
A cutting point between x abd x' is w[x] + w[x']. 
This can be modelled as a single number, problem into k elems in a heap kek.
God forbid you see this one in an interview.

"""


class Solution:
    def putMarbles(self, weights: List[int], k: int) -> int:
        if k == 1:
            return 0

        cuttingPoint = [weights[i] + weights[i + 1] for i in range(len(weights) - 1)]
        cuttingPoint.sort()

        # Sum of last K elem - Sum of first K elem
        # index[-(k-1):0]  and index[:k-1]
        return sum(cuttingPoint[-(k - 1) :]) - sum(cuttingPoint[: k - 1])
