import heapq
"""

Observation: get all the consecutive baloons, remove all the smallest

"""

class Solution:
    def minCost(self, colors: str, neededTime: List[int]) -> int:
        res = 0 
        i = 0 
        while i < len(colors):
            color = colors[i]
            j = i+1
            cost = [neededTime[i]]
            while  j < len(colors) and color == colors[j]:
                cost.append(neededTime[j])
                j+=1

            heapq.heapify(cost)
            while len(cost) >1:
                c = heapq.heappop(cost)
                res += c

            # Move first pointer to second pointer
            i = j

        return res
