from typing import List

"""

Question 2747: https://leetcode.com/problems/count-zero-request-servers/

Constraints:
1 <= n <= 1e5
1 <= logs.length <= 1e5
1 <= queries.length <= 1e5
logs[i].length == 2
1 <= logs[i][0] <= n
1 <= logs[i][1] <= 1e6
1 <= x <= 1e5
x < queries[i] <= 1e6

Idea: 
- sort both the logs and the queries. store sortedQueries into a separate arrays.
- iterate over queries, try to increase logs to before ends of a query.
"""


class Solution:
    def countServers(
        self, n: int, logs: List[List[int]], x: int, queries: List[int]
    ) -> List[int]:

        logs.sort(key=lambda x: x[1])
        queriesSorted = sorted(queries)

        resultMap = dict()  # K: Queries value, V: inactive server count
        activeServer = dict()  # K: serverID, V: most recent active time
        res = []
        logIndexStart, logIndexEnd, queriesIndex = 0, 0, 0
        for end in queriesSorted:
            start = end - x
            while logIndexStart < len(logs) and logs[logIndexStart][1] <= end:
                request = logs[logIndexStart]
                if request[1] >= start:
                    activeServer[request[0]] = request[1]
                logIndexStart += 1

            while logIndexEnd < len(logs) and logs[logIndexEnd][1] < start:
                request = logs[logIndexEnd]
                if activeServer.get(request[0], -1) == request[1]:
                    activeServer.pop(request[0])
                logIndexEnd += 1

            if queriesIndex < len(queriesSorted) and end == queriesSorted[queriesIndex]:
                resultMap[queriesSorted[queriesIndex]] = n - len(activeServer)
                queriesIndex += 1

        for q in queries:
            res.append(resultMap[q])

        return res
