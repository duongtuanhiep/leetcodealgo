from typing import List

"""

Question 1007: https://leetcode.com/problems/minimum-domino-rotations-for-equal-row

Approach:
- try to fit 1->6 to both tops and bottom. Use cnt to eliminate possible pitfall.

- we can only fit if theres at least cnt[x] == len(arr). At most there will be 2. Since theres only 2 values, theres ounly 1 so we only need to try swap one.
- For each of it counts elem from bottom and top.

"""


class Solution:
    def minDominoRotations(self, tops: List[int], bottoms: List[int]) -> int:
        cnt = [0] * 6
        for i in range(len(tops)):
            cnt[tops[i] - 1] += 1
            if tops[i] != bottoms[i]:
                cnt[bottoms[i] - 1] += 1

        target = -1
        for i, count in enumerate(cnt):
            if count == len(tops):
                target = i + 1

        if target == -1:
            return target
        return len(tops) - max(tops.count(target), bottoms.count(target))
