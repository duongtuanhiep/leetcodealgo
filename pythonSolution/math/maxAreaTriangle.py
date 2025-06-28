from typing import List
from collections import defaultdict


"""

Question 3588: https://leetcode.com/problems/find-maximum-area-of-a-triangle/

Idea:
- Create an adjacency list of all the points on X, Y axies
- 
"""


class Solution:
    def maxArea(self, coords: List[List[int]]) -> int:
        xaxis, yaxis = defaultdict(list), defaultdict(list)

        for coord in coords:
            xaxis[coord[0]].append(coord[1])
            yaxis[coord[1]].append(coord[0])

        flattenx = sorted([xcoord for xcoord in xaxis.keys()])
        flatteny = sorted([ycoord for ycoord in yaxis.keys()])

        for x, points in xaxis.items():
            xaxis[x] = sorted(points)

        for y, points in yaxis.items():
            yaxis[y] = sorted(points)

        # Get the furthest points from x,y
        res = 0
        for x, points in xaxis.items():
            if len(points) > 1:
                # Find the best remaining points according to x axis
                width = points[-1] - points[0]
                height = max(x - flattenx[0], flattenx[-1] - x)
                res = max(res, width * height)

        for y, points in yaxis.items():
            if len(points) > 1:
                # Find the best remaining points according to x axis
                width = points[-1] - points[0]
                height = max(y - flatteny[0], flatteny[-1] - y)
                res = max(res, width * height)

        return -1 if not res else res
