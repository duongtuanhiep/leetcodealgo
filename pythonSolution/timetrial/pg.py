from typing import List
import heapq


class Solution:
    def networkDelayTime(self, times: List[List[int]], n: int, k: int) -> int:
        # Create arrays that holds min_cost[node number] = cost_to_node
        min_cost = [1e9 if i != k else 0 for i in range(n+1)] 
        
        # Create arrays that holds possible directions[node number] = [(destination node, weight)]
        graphs = [[] for _ in range(n+1)] 

        # Populate the graphs
        for edge in times:
            graphs[edge[0]].append([edge[1], edge[2]])

        # create an array that later will be used as min heap / prio queue
        pq = [(0, k)]

        # Start BFS from the node that has the smallest cost 
        while pq:
            cur_cost, cur_node = heapq.heappop(pq)

            # This is a non optimal path since previous is already in min_cost
            if (
                cur_cost > min_cost[cur_node]
            ):  
                continue

            # From current node, try calculate new distance to all neighboring nodes
            for next_node, cost in graphs[cur_node]:
                next_cost = cost + cur_cost

                # If next distance is smaller, add it to the min heap
                if next_cost < min_cost[next_node]:
                    min_cost[next_node] = next_cost
                    heapq.heappush(pq, (next_cost, next_node))

        res = 0
        for cost in min_cost[1:]:
            if cost > res:
                res = cost

        if res == 1e9:
            return -1
        return res
