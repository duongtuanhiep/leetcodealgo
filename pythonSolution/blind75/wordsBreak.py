class Solution:
    def wordBreak(self, s: str, wordDict: List[str]) -> bool:
        available = set(wordDict)
        dp = [0 for i in range(len(s)+1)]
        return self.canFit(s, 0, dp, available)

    # start = 0, end is len
    def canFit(self, s: str, start: int, dp: List[int], available: set[str]) -> bool:
        if start == len(s):
            return True
        elif start == len(s)-1:
            return s[start:] in available

        for i in range(start, len(s)+1):
            if s[start:i] in available and dp[i] == 0:
                canFitLess = self.canFit(s, i, dp , available)
                if canFitLess:
                    return True
                dp[i] = 1

        return False

