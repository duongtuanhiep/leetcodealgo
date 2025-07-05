from typing import List
from collections import deque, defaultdict


class Solution:
    def alienOrder(self, words: List[str]) -> str:
        """
        Most intuitive way is probably try to reverse the lexical by recursions
        if we build stuff out at worst we'll have 1e5 nodes which is fine ?

        One solutions is to build adj list based off character.
        """

        adjList = dict()
        checked = set()
        buildAdjList(adjList, checked, words, "")

        checked = dict()
        res = deque()
        cyclic = Dfs(adjList, checked, res, "")
        if cyclic:
            return ""

        return "".join(res)


def buildAdjList(adjList, checked, words, prefix):
    if prefix in checked:
        return
    checked.add(prefix)

    prefLen = len(prefix)
    parent = ""
    for word in words:
        if prefLen >= len(word):
            continue

        char = word[prefLen]

        # Build a adjacency list of all char in the same index if has same prefix
        if not prefix or prefix == word[:prefLen]:
            if parent not in adjList:
                adjList[parent] = set()
            adjList[parent].add(char)
            parent = char

            nextPrefix = prefix + char

            if nextPrefix not in checked:
                buildAdjList(adjList, checked, words, nextPrefix)


def Dfs(adjList, checked, res, char) -> bool:
    if char in checked:
        return False
    checked[char] = 1

    if char not in adjList:
        checked[char] = 2
        res.appendleft(char)
        return False

    for nextChar in adjList[char]:
        if nextChar == char:
            continue

        if checked.get(nextChar, 0) == 1:
            return True

        if nextChar not in checked and nextChar != char:
            if Dfs(adjList, checked, res, nextChar):
                return True

    checked[char] = 2
    res.appendleft(char)
    return False


if __name__ == "__main__":
    cases = [
        ["wrt", "wrf", "er", "ett", "rftt"],
        ["z", "x"],
        ["z", "x", "z"],
    ]

    for c in cases:
        print(Solution().alienOrder(c))
