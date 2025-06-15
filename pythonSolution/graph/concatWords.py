from typing import List


"""

Question 472: https://leetcode.com/problems/concatenated-words

Either build a Trie like a weakling or try to dfs + memo on character like a big boy.
Needs to cache previous function run to avoid extra computation on inputs like: 
["a","aa","aaa",...,"aaaaaaaaaaaaaaa"]

Cheatcode: @lru_cache(None) on normal DFS, can easily simulate this function by using cache[$functionInput] = $functionOutput
"""


class Solution:
    def findAllConcatenatedWordsInADict(self, words: List[str]) -> List[str]:
        exist = set()
        for w in words:
            exist.add(w)

        lruCache = dict()

        # @lru_cache(None) - Cheatcode if lazy
        def canBeBrokenDown(word: str) -> int:
            nonlocal exist, lruCache
            if not word:
                return 0
            if word in lruCache:
                return lruCache[word]

            currentWord = ""
            partCnt = -1
            for i, char in enumerate(word):
                currentWord += char
                if currentWord in exist:
                    remain = word[i + 1 :]
                    remainCnt = canBeBrokenDown(remain)
                    if remainCnt != -1:
                        partCnt = max(remainCnt + 1, partCnt)

            lruCache[word] = partCnt

            return partCnt

        return [word for word in words if canBeBrokenDown(word) > 1]
