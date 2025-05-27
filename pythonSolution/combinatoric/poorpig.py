"""

Question 458: https://leetcode.com/problems/poor-pigs/

Insane encoding problem.
Reference:
- Hint
- https://leetcode.com/problems/poor-pigs/description/comments/1566889/

"""


class Solution:
    def poorPigs(self, buckets: int, minutesToDie: int, minutesToTest: int) -> int:
        pigCnt = 0
        tries = minutesToTest // minutesToDie
        while (tries + 1) ** pigCnt < buckets:
            pigCnt += 1

        return pigCnt
