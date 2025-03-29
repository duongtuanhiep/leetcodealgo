import heapq
from typing import List

"""

Question 2594: https://leetcode.com/problems/minimum-time-to-repair-cars

Solution: 
- Bucketing the mechanic rank since it's [1:100]
- Put the next repair time for n+1 cars into a heap
- Iterating through the heap, for each element increase the car count of current mechanic groups

"""

class Solution:
    def repairCars(self, ranks: List[int], cars: int) -> int:
        rank_bucket = [0 for _ in range(101)]
        for r in ranks:
            rank_bucket[r] += 1

        # Make a heap to pick the best candidate groups
        pq = []

        # Init repair time and put in heap (if each mechanic repair 1 car)
        current_count = [0 for _ in range(101)]
        for rank, count in enumerate(rank_bucket):
            if count:
                # heapitems (nextCarCost, number of cars)
                heapq.heappush(pq, (rank, rank))

        # Simulate the car repairing by distributing to the lowest time mechanic
        while cars > 0:
            mechanic = heapq.heappop(pq)
            current_count[mechanic[1]] += 1
            new_time = (mechanic[1]) * (current_count[mechanic[1]] + 1) ** 2
            heapq.heappush(pq, (new_time, mechanic[1]))

            res = mechanic[0]
            cars -= rank_bucket[mechanic[1]]

        return res

cases = [
    [[4,2,3,1],10], # 16
    [[5,1,8],6], # 16
    [[99,1,1,1,1,1,1],9],
    [[3,2,5,5,6,7,3,3,2,4,5,6,3,4,6,3,7,54,32,12,23,69],100], #128
]

for c in cases:
    print(Solution().repairCars(c[0], c[1]))
