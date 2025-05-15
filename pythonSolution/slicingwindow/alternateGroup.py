from typing import List

"""

Question 3208: https://leetcode.com/problems/alternating-groups-ii

Observation: 
- Index can be the "middle" (not start/end) when theres "010" or "101"
- We can mark "middle" index as true and others index as false. Previous arrs would be "FTF" and "FTF"
- An alternating colors range would be any range that statisfy "(T|F)T..(T|F)"

Solution: Sliding windows
- Make colors 

"""


class Solution:
    def numberOfAlternatingGroups(self, colors: List[int], k: int) -> int:
        alt = [False for _ in colors] # Store if index i can be a "middle" or not
        
        for i, val in enumerate(colors):
            left = i - 1
            right = i+1 if i+1<len(colors) else (i+1)%len(colors)
            if val != colors[right] and val != colors[left]:
                alt[i] = True
            else:
                alt[i] = False

        res = 0
        res_arr = list()
        i = 0 

        while  0 <= i < len(colors):
            if not alt[i]:
                j = i+1 if i+1 < len(alt) else (i+1)%len(alt)
                step = 1
                while j != i:
                    if alt[j]:
                        step +=1
                        j = j+1 if j+1 < len(alt) else (j+1)%len(alt)
                    else:
                        step +=1
                        res_arr.append(step)
                        break
            i+=1

        for step in res_arr:
            if step >=k:
                res += step - k + 1

        if len(res_arr) == 0: # only happens when theres no False in alt
            return len(alt)

        return res

if __name__ == "__main__":
    cases = [
        # [[0,1,0,1,0], 3], # 3
        # [[0,1,0,0,1,0,1],6], # 2
        # [[1,1,0,1],4], # 0
        [[0,1,0,1],3] # 2
    ]

    for c in cases:
        print(Solution().numberOfAlternatingGroups(c[0],c[1]))
