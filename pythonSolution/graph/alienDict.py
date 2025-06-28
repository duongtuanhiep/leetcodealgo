"""

Leetcode 269: https://leetcode.com/problems/alien-dictionary/description/?envType=company&envId=amazon&favoriteSlug=amazon-six-months

"""


class Solution:
    def alienOrder(self, words: List[str]) -> str:
        """
        2 Important steps:
        From the input words, try to build a graph that later can be sorted via topo
        The ways to do it is to compare the letter[i] of words[a] and words[b] if they have same prefix
        If the condition above is sastified then we iteratively cycle through all of them and make a
        "linked list" with currentNode from previous steps as parents
        -> Watch out for lexicographically order problems

        From the Adjlist that has K:V of Parent:[]Child, we perform a toposort with cyclic detection.
        """

        adjList = dict()
        checked = set()
        if buildAdjList(adjList, checked, words, "", 100):
            return ""

        checked = dict()
        res = deque()
        if Dfs(adjList, checked, res, ""):
            return ""

        return "".join(res)


def buildAdjList(adjList, checked, words, prefix, indexOfPrefix) -> bool:
    if prefix in checked:
        return
    checked.add(prefix)

    prefLen = len(prefix)
    parent = ""
    for i, word in enumerate(words):
        # We do 2 things here: checks if it's worth exploring and checks if input is bs.
        if prefLen >= len(word):
            # This is for the test cases ["z", "z"] -> "z" and ["abc","ab"] -> ""
            if (
                prefix == word
                and indexOfPrefix < 100
                and words[i] != words[indexOfPrefix]
                and i != indexOfPrefix
            ):
                return True
            continue

        char = word[prefLen]

        # Build a adjacency list of all char in the same index if has same prefix
        if not prefix or prefix == word[:prefLen]:
            if parent not in adjList:
                adjList[parent] = set()

            # "cycle" through all matching prefix words, keep adding char as child of previous char
            adjList[parent].add(char)
            parent = char

            # Continue down with DFS
            nextPrefix = prefix + char
            if nextPrefix not in checked:
                # Propagate cyclic information here
                if buildAdjList(adjList, checked, words, nextPrefix, i):
                    return True

    return False


def Dfs(adjList, checked, res, char) -> bool:
    if char in checked:
        return False

    # Consider current items as in progress
    checked[char] = 1

    if char not in adjList:
        # If char aint found we can confidently mark it early and move on
        checked[char] = 2
        res.appendleft(char)
        return False

    for nextChar in adjList[char]:
        # This is fine cause unlike normal topo it can be linked here
        if nextChar == char:
            continue

        # Cyclic checks
        if checked.get(nextChar, 0) == 1:
            return True

        # Only explore new Node
        if nextChar not in checked:
            # Propagate cyclic detection
            if Dfs(adjList, checked, res, nextChar):
                return True

    checked[char] = 2
    res.appendleft(char)
    return False
