from typing import List

"""

Question 683: https://leetcode.com/problems/k-empty-slots

Confusing questions
"""


class FenwickTree:
    def __init__(self, size):
        self.tree = [0] * (size + 1)

    def update(self, index, delta):
        index += 1
        while index < len(self.tree):
            self.tree[index] += delta
            index += index & -index

    def query(self, index) -> int:
        index += 1
        res = 0
        while index > 0:
            res += self.tree[index]
            index -= index & -index
        return res


class Solution:
    def kEmptySlots(self, bulbs: List[int], k: int) -> int:
        if len(bulbs) == 1:
            return -1

        n, distance = len(bulbs), k + 1

        reverseMapping = dict()
        for index, position in enumerate(bulbs):
            reverseMapping[position] = index + 1

        fwtree = FenwickTree(n + 1)
        for turnOnTime in range(1, n + 1):
            lightNumber = bulbs[turnOnTime - 1]
            fwtree.update(lightNumber, 1)
            lightPosition = reverseMapping[bulbs[turnOnTime - 1]]
            if lightPosition == 1:
                continue
            previousBulb, nextBulb = lightNumber - distance, lightNumber + distance

            if nextBulb <= n and reverseMapping[nextBulb] < turnOnTime:
                if fwtree.query(nextBulb) - fwtree.query(lightNumber) == 1:
                    return turnOnTime

            if previousBulb > 0 and reverseMapping[previousBulb] < turnOnTime:
                if fwtree.query(lightNumber) - fwtree.query(previousBulb) == 1:
                    return turnOnTime

        return -1
