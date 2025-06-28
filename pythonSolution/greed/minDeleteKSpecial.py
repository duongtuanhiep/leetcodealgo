from collections import Counter

"""

Questions 3085: https://leetcode.com/problems/minimum-deletions-to-make-string-k-special

Idea:
- Counts the frequencies of each char then sort it. 
- Out put would be something like [1,3,6,34,42,54]
- for each items in the range, we calculate the number of deletion needed to sastify it.
"""


class Solution:
    def minimumDeletions(self, word: str, k: int) -> int:
        freqArr = [val for val in Counter(word).values()]
        freqArr.sort()

        deletion = len(word)
        for freq in freqArr:
            deletion = min(deletion, deleteCount(freq, freq + k, freqArr))

        return deletion


def deleteCount(minCount, maxCount, freqArr) -> int:
    res = 0
    for cnt in freqArr:
        if cnt < minCount:
            res += cnt
        elif cnt > maxCount:
            res += cnt - maxCount
    return res
