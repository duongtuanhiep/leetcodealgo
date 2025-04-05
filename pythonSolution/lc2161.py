"""

Question 2161: https://leetcode.com/problems/partition-array-according-to-given-pivot/

Dumb solutions: (3 passes)
- build 3 arrs for number that is smaller, equal and greater and concatenate them.
"""


class Solution:
    def pivotArray(self, nums: List[int], pivot: int) -> List[int]:
        lower = [x for x in nums if x < pivot]
        equal = [x for x in nums if x == pivot]
        upper = [x for x in nums if x > pivot]
        return lower + equal + upper
