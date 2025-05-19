from typing import List

"""

Question 2900: https://leetcode.com/problems/longest-unequal-adjacent-groups-subsequence-i

Greadily try to build both sequence, one start with 0 and one start with 1
"""


class Solution:
    def getLongestSubsequence(self, words: List[str], groups: List[int]) -> List[str]:
        start0, start1 = [], []
        for i, s in enumerate(words):
            if groups[i] == 0 and not start0:
                start0.append(i)
            if groups[i] == 1 and not start1:
                start1.append(i)

            if start0 and groups[start0[-1]] != groups[i]:
                start0.append(i)
            if start1 and groups[start1[-1]] != groups[i]:
                start1.append(i)

        res = start0 if len(start0) > len(start1) else start1

        return [words[i] for i in res]
