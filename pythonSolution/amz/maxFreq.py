from typing import List
from collections import Counter

"""

Question 3434: https://leetcode.com/problems/maximum-frequency-after-subarray-operation

"""


"""
This didn't work, cause [28,15,15,15,28,40,23,23,23,23] / K = 28 is 6
"""


class Solution:
    def maxFrequency(self, nums: List[int], k: int) -> int:
        freq = [0] * 51  # windows range
        startIndex, startFreq, cnt = 0, nums.count(k), 0
        # drop the "accumulating if it goes below 0"
        for i, n in enumerate(nums):
            freq[n] += 1
            if n != k and max(freq) <= freq[k]:
                freq = [0] * 51

            cnt = max(cnt, startFreq + freq[n] - freq[k])

        return cnt

    def maxFrequency2(self, nums: List[int], k: int) -> int:
        cnt = Counter()
        res = 0
        for n in nums:
            cnt[n] = max(cnt[n], cnt[k]) + 1
            res = max(res, cnt[n] - cnt[k])
        return res + cnt[k]
