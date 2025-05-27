"""
Question 2131: https://leetcode.com/problems/longest-palindrome-by-concatenating-two-letter-words
"""


class Solution:
    def longestPalindrome(self, words: List[str]) -> int:
        wordCnt = dict()
        for w in words:
            wordCnt[w] = wordCnt.get(w, 0) + 1

        res = 0
        hasMid = False
        visited = set()
        # Try to greadily put words
        for word, cnt in wordCnt.items():
            reversedWord = word[1] + word[0]

            if word in visited or reversedWord in visited:
                continue
            visited.add(word)
            visited.add(reversedWord)

            if reversedWord == word:
                if not hasMid and cnt % 2 == 1:  # We try to take odd count here
                    hasMid = True
                    res += cnt * 2
                elif cnt >= 2:
                    res += (cnt // 2) * 4
            elif wordCnt.get(reversedWord, 0):
                res += min(cnt, wordCnt[reversedWord]) * 4

        return int(res)
