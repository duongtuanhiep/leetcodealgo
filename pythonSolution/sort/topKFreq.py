from typing import List
from collections import Counter

"""

Question 347: https://leetcode.com/problems/top-k-frequent-elements

"""


class Solution:

    # O(NLogK) solutions with k as no of uniqueElems
    def topKFrequent2(self, nums: List[int], k: int) -> List[int]:
        return [count[0] for count in Counter(nums).most_common(k)]

    # O(N) with BucketSort
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        cnt = Counter(nums)
        # maximum N+1 buckets since it's capped by frequency
        buckets = [[] for n in range(len(nums) + 1)]

        for num, freq in cnt.items():
            buckets[freq].append(num)

        res = []
        for i in range(len(buckets) - 1, 0, -1):
            if buckets[i]:
                for n in buckets[i]:
                    res.append(n)

        return res[:k]

    # O(N) average with quick selects
