"""
Question 838: https://leetcode.com/problems/push-dominoes

Solutions
2 pointers - O(N): Quite nasty since you have to account for cases manually.

Alternatively, https://leetcode.com/problems/push-dominoes
This is considered auxilary dp (for whats it worth)
Running 2 passes to calculate distance for each cell to nearest L and R.
From that we can built the result.
for i in dominoes:
    i == "L" if dist_L[i] < dist_R[i]
    i == "R" if dist_L[i] > dist_R[i]
    i == "."
"""


class Solution:
    def pushDominoes(self, dominoes: str) -> str:
        res, i = "", 0
        while i < len(dominoes):
            # keep extend till we see "L"
            firstRightIndex = -1
            rightIndex = -1
            leftIndex = -1
            j = i
            while j < len(dominoes):
                # Store first R occurence
                if dominoes[j] == "R" and rightIndex == -1:
                    firstRightIndex = j
                if dominoes[j] == "R":
                    rightIndex = j
                if dominoes[j] == "L":
                    leftIndex = j
                    j += 1
                    break
                j += 1

            print(rightIndex, leftIndex)
            if leftIndex >= 0 and rightIndex >= 0:
                # Case: R..L, R...L, R.L
                dist = leftIndex - rightIndex - 1
                for k in range(i, leftIndex + 1):
                    if firstRightIndex <= k <= rightIndex + dist // 2:
                        res += "R"
                    elif leftIndex >= k >= leftIndex - dist // 2:
                        res += "L"
                    else:
                        res += "."
            elif leftIndex >= 0:
                for _ in range(i, leftIndex + 1):
                    res += "L"
            elif firstRightIndex >= 0:
                for k in range(i, j):
                    if k <= firstRightIndex:
                        res += "."
                    else:
                        res += "R"
            else:
                for _ in range(i, j):
                    res += "."

            i = j

        return res
