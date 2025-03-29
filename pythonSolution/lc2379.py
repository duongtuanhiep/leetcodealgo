
"""

Question 2379: https://leetcode.com/problems/minimum-recolors-to-get-k-consecutive-black-blocks

Solution: Sliding windows
- Start with i,j = 0
- move j until theres k occurence of 'B'
- Counts 'W', put to res if count < res || count > -1
- move i forward, adjust j according to i
- return when j == len()


"""

class Solution:
    def minimumRecolors(self, blocks: str, k: int) -> int:
        i,j,b_count,res,w_count =0,0,0,-1,0
        while 0<= i<=j < len(blocks):
            while j-i < k and j < len(blocks):
                if blocks[j] == "W":
                    w_count +=1
                if blocks[j] == "B":
                    b_count +=1
                j+=1

            if w_count < res or res == -1: 
                res = w_count
    
            if i + k <len(blocks):
                if blocks[i] == "B":
                    b_count -=1
                else:
                    w_count -=1
                i+=1
            else: 
                break
            
        return res

if __name__ == "__main__":
    cases = [
        ["WBBWWBBWBW",7],
        ["WBWBBBW",3],
        ["WWWBWWBWBWBWWBBBWWWBBW",5]
    ]

    for c in cases:
        print(Solution().minimumRecolors(c[0],c[1]))
