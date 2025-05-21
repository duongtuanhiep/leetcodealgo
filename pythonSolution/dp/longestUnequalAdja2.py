from typing import List
from collections import deque
"""

Question 2901: https://leetcode.com/problems/longest-unequal-adjacent-groups-subsequence-ii/


Could be solved by DP or Kahn algorithm, turning this into topo sort problems
"""


class Solution:
    def getWordsInLongestSubsequence(self, words: List[str], groups: List[int]) -> List[str]:
        # Build a hamming dist where hd[i]=[[j],[k]]
        hammingDst = [[] for _ in range(len(words))] 
        for i, w in enumerate(words):
            for j, w2 in enumerate(words):
                if i == j or len(w) != len(w2):
                    continue
                if self.hammingDist(w, w2) == 1:
                    hammingDst[i].append(j)
                    
        # Building dp - this store previous indexes and current length.
        dp = [() for _ in range(len(words))] 
        for i, w in enumerate(words):
            candidate = []
            for k in hammingDst[i]:
                if k < i and groups[k] != groups[i]: 
                    dist,last = dp[k][0], k
                    candidate.append((dist, k))
            if not candidate:
                dp[i] = (1, -1)
            else:
                best = max(candidate)
                dp[i] = (best[0]+1, best[1])
        
        best = max(dp)
        firstIndex = dp.index(best)
        resArr = deque()
        resArr.appendleft(firstIndex)
        while best[1] != -1:
            resArr.appendleft(best[1])
            best = dp[best[1]]

        return [words[i] for i in resArr]

    def hammingDist(self, a: str, b: str) -> int:
        ans = 0
        for i in range(len(a)):
            if a[i] != b[i]:
                ans += 1

        return ans
