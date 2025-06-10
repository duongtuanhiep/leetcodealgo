from typing import List
from collections import Counter

"""

Question 347: https://leetcode.com/problems/top-k-frequent-elements

"""


class Solution:

    # O(NLogK) solutions with k as no of uniqueElems
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        return [count[0] for count in Counter(nums).most_common(k)]

    # O(N) Average solutions
