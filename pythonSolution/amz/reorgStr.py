import heapq

"""

Question 767: https://leetcode.com/problems/reorganize-string

"""


class Solution:
    def reorganizeString(self, s: str) -> str:
        freq = [0] * 26
        for char in s:
            pos = ord(char) - ord("a")
            freq[pos] += 1

        maxLetter = []
        for pos, f in enumerate(freq):
            if f != 0:
                char = chr(pos + ord("a"))
                heapq.heappush(maxLetter, (-f, char))

        s = ""
        while maxLetter:
            count, optionChar = heapq.heappop(maxLetter)
            if s and s[-1] == optionChar:
                if not maxLetter:
                    return ""
                newCnt, possibleChar = heapq.heappop(maxLetter)
                if newCnt == 0:
                    return ""
                s += possibleChar
                heapq.heappush(maxLetter, (count, optionChar))
                if newCnt + 1 != 0:
                    heapq.heappush(maxLetter, (newCnt + 1, possibleChar))
            else:
                s += optionChar
                if count + 1 != 0:
                    heapq.heappush(maxLetter, (count + 1, optionChar))
        return s


if __name__ == "__main__":
    cases = [
        "aba",
        "aaabba",
    ]

    for c in cases:
        print(Solution().reorganizeString(c))
