from typing import List

"""

Question 2523: https://leetcode.com/problems/closest-prime-numbers-in-range

"""

class Solution:
    def closestPrimes(self, left: int, right: int) -> List[int]:
        sieves=[True for _ in range(right+1)]
        sieves[0], sieves[1]= False, False


        prime = list()
        for i in range(2, right+1):
            if sieves[i] and \
                (i % 2 ==1 or i == 2):
                if left <= i <= right:
                    prime.append(i)
                for j in range(i*i, right+1, i):
                    sieves[j]=False
        
        minGap, res = -1,[-1,-1]
        for i in range(len(prime)):
            if i+1 < len(prime): 
                gap = prime[i+1] - prime[i]
                if minGap == -1 or gap < minGap:
                    minGap = gap
                    res = [prime[i], prime[i+1]]

        return res

if __name__ == "__main__":
    cases = [
        [10, 19],
        [4,6],
        [19,31]
    ]

    solution = Solution()

    for case in cases:
        print(solution.closestPrimes(case[0], case[1]))
