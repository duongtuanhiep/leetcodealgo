from typing import List
import heapq

"""

Question: 3542: https://leetcode.com/problems/minimum-operations-to-convert-all-elements-to-zero

This question can be solved greedily and there is a few ways: Monotonic stack or segment trees
Bellow implementation is heap+set which function like monotonic stack but worse.

One sample: https://leetcode.com/problems/minimum-operations-to-convert-all-elements-to-zero/solutions/6731695/100-beat-monotonic-increasing-stack-o-n-or-segment-tree-o-n-logn-explained/
"""

class Solution:
    def minOperations(self, nums: List[int]) -> int:
        ans = 0
        maxq = []
        seen = set()
        for n in nums:
            while len(maxq) > 0 and n < -maxq[0]:
                num = -heapq.heappop(maxq)
                if num in seen: 
                    seen.remove(num)
            
            if n != 0 and (len(maxq) == 0 or n >= -maxq[0]):
                if n not in seen:
                    seen.add(n)
                    ans +=1
                heapq.heappush(maxq, -n)

        return ans
