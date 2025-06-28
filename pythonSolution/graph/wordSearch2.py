from typing import List

"""

Question 212: https://leetcode.com/problems/word-search-ii

"""


class Solution:
    def findWords(self, board: List[List[str]], words: List[str]) -> List[str]:
        directions = [(1, 0), (0, 1), (-1, 0), (0, -1)]
        n, m = len(board), len(board[0])

        cache = dict()

        def canFind(word: str, i, j, step, visited) -> bool:
            nonlocal directions, board, m, n

            visited.add((i, j))
            step += 1
            if step == len(word):
                return True

            for dI, dJ in directions:
                nextI, nextJ = i + dI, j + dJ
                if 0 <= nextI < n and 0 <= nextJ < m:
                    if (
                        board[nextI][nextJ] == word[step]
                        and (nextI, nextJ) not in visited
                    ):
                        validPath = canFind(word, nextI, nextJ, step, visited)
                        if validPath:
                            return True

            visited.discard((i, j))

            return False

        res = []
        for w in words:
            found = False
            for i in range(n):
                for j in range(m):
                    if w[0] == board[i][j] and canFind(w, i, j, 0, set()):
                        found = True
                        res.append(w)
                        break
                    if found:
                        break
                if found:
                    break

        return res


class Solution:
    def findWords(self, board: List[List[str]], words: List[str]) -> List[str]:
        directions = [(1, 0), (0, 1), (-1, 0), (0, -1)]
        n, m = len(board), len(board[0])

        def canFind(word: str, i, j) -> bool:
            nonlocal directions, board, m, n
            q = [(i, j, 0)]
            visited = set()
            while q:
                curI, curJ, step = q.pop()
                visited.add((curI, curJ))
                step += 1
                if step == len(word):
                    return True

                hasNext = False
                for dI, dJ in directions:
                    nextI, nextJ = curI + dI, curJ + dJ
                    if 0 <= nextI < n and 0 <= nextJ < m:
                        if (
                            board[nextI][nextJ] == word[step]
                            and (nextI, nextJ) not in visited
                        ):
                            q.append((nextI, nextJ, step))
                            hasNext = True

                if not hasNext:
                    visited.discard((curI, curJ))

            return False

        res = []
        for w in words:
            found = False
            for i in range(n):
                for j in range(m):
                    if w[0] == board[i][j] and canFind(w, i, j):
                        found = True
                        res.append(w)
                        break
                    if found:
                        break
                if found:
                    break

        return res
