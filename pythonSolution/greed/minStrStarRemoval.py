from collections import defaultdict

"""

Question 3170: https://leetcode.com/problems/lexicographically-minimum-string-after-removing-stars

"""


class Solution:
    def clearStars(self, s: str) -> str:
        charCnt = defaultdict(list)
        pick = [True for _ in range(len(s))]
        for i, char in enumerate(s):
            if char == "*":
                lowestChar = min(charCnt)
                charPos = charCnt[lowestChar]
                deleted = charPos.pop()
                pick[deleted] = False
                if not charPos:
                    charCnt.pop(lowestChar)
            else:
                charCnt[char].append(i)

        return "".join([s[i] for i in range(len(s)) if pick[i] and s[i] != "*"])
