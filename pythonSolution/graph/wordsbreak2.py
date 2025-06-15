from typing import List
from collections import defaultdict

"""

Question 140: https://leetcode.com/problems/word-break-ii/

Solutions:
- Perform DFS on the input with givens words. We can put it in a set for O(1) checks
- If you feeling fancy, build a Trie and traverse that graph instead. 

"""


class TrieNode:
    def __init__(self):
        self.children = {}
        self.isEnd = False


class Trie:
    def __init__(self):
        self.root = TrieNode()

    def add(self, word):
        node = self.root
        for letter in word:
            if letter not in node.children:
                node.children[letter] = TrieNode()
            node = node.children[letter]
        node.isEnd = True

    def search(self, word) -> bool:
        node = self.root
        for letter in word:
            if letter not in node.children:
                return False
            node = node.children[letter]
        return node.isEnd


class Solution:
    def wordBreak(self, s: str, wordDict: List[str]) -> List[str]:
        memo = defaultdict(list)

        # Build Trie
        trie = Trie()
        for w in wordDict:
            trie.add(w)

        #  @lru_cache(None)  # memoization to avoid recomputation
        def memoDfs(current: str) -> list[str]:
            nonlocal memo, trie
            if current in memo:
                return memo[current]

            currStr = ""
            node = trie.root
            res = []
            for i, letter in enumerate(current):
                currStr += letter
                if letter not in node.children:
                    break

                node = node.children[letter]
                if not node.isEnd:
                    continue

                remainingStr = current[i + 1 :]
                combinations = memoDfs(remainingStr)
                if not remainingStr:
                    res.append(currStr)  # Only one cases
                else:
                    for c in combinations:
                        res.append(currStr + " " + c)

            memo[current] = res
            return res

        return memoDfs(s)
