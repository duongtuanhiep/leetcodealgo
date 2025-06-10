from typing import List
from collections import defaultdict

"""

Question 1152: https://leetcode.com/problems/analyze-user-website-visit-pattern

"""


class Solution:
    def mostVisitedPattern(
        self, username: List[str], timestamp: List[int], website: List[str]
    ) -> List[str]:
        allPatternCnt = dict()
        userVisits = defaultdict(list)
        # Bundle into users

        logs = sorted(zip(username, timestamp, website))
        for user, _, site in logs:
            userVisits[user].append(site)

        for visitHistory in userVisits.values():
            n = len(visitHistory)
            if n < 3:
                continue

            userPatterns = set()
            for i in range(n):
                for j in range(i + 1, n):
                    for k in range(j + 1, n):
                        userPatterns.add(
                            (visitHistory[i], visitHistory[j], visitHistory[k])
                        )

            for pattern in userPatterns:
                allPatternCnt[pattern] = allPatternCnt.get(pattern, 0) + 1

        bestFreq = allPatternCnt[max(allPatternCnt, key=allPatternCnt.get)]

        return min(
            [
                pattern
                for pattern in allPatternCnt.keys()
                if allPatternCnt[pattern] == bestFreq
            ]
        )
