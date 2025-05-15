from typing import List
from collections import deque

"""

Question 862: https://leetcode.com/problems/shortest-subarray-with-sum-at-least-k


Idea: 
- Keep track of indexes and it's prefix sum. The result will be i - some_prev_index + 1
- Maintain deque consisting of only increasing prefix sum. This is curcial.

Solutions inspired from: https://leetcode.com/problems/shortest-subarray-with-sum-at-least-k/solutions/143726/c-java-python-o-n-using-deque/?envType=problem-list-v2&envId=monotonic-queue

"""


class Solution:
    def shortestSubarray(self, nums: List[int], k: int) -> int:
        res = -1
        rollingSum = 0
        q = deque() # Store rolling sum / index
        q.append((0,0))
        for i, n in enumerate(nums):
            rollingSum += n
            
            # This part is similar to only positive sliding windows
            while q and rollingSum - q[0][0] >= k:
                left = q.popleft()
                if res == -1 or res > i - left[1]+1:
                    res = i - left[1] + 1

            # Maintaining monocity
            while q and rollingSum < q[-1][0]:
                q.pop()

            q.append((rollingSum,i+1)) 

        return res




if __name__ == "__main__":
    cases = [
        [[1, 2], 4],  # -1
        [[2, -1, 2], 3],  # 3
        [[17, 85, 93, -45, -21], 150],  # 2
        [[84,-37,32,40,95], 167], # 3
        [[-34,37,51,3,-12,-50,51,100,-47,99,34,14,-13,89,31,-14,-44,23,-38,6], 151], # 2
    ]

    for c in cases:
        print(Solution().shortestSubarray(c[0],c[1]))



"""
Magic
    def shortestSubarray(self, A, K):
        d = collections.deque([[0, 0]])
        res, cur = float('inf'), 0
        for i, a in enumerate(A):
            cur += a
            while d and cur - d[0][1] >= K:
                res = min(res, i + 1 - d.popleft()[0])
            while d and cur <= d[-1][1]:
                d.pop()
            d.append([i + 1, cur])
        return res if res < float('inf') else -1

"""
