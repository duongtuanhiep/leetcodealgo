
"""

Question 647: https://leetcode.com/problems/palindromic-substrings/?envType=problem-list-v2&envId=oizxjoit

Optimal solution achieved by Manacher Algoritm
"""


class Solution:
    def countSubstrings(self, s: str) -> int:
        res = 0

        for i in range(len(s)):
            res += 1 # Single char is also a palindrome
            if i +1<len(s) and s[i] == s[i+1]:  # check palindrome of "2" center
                res+=1
                j,k = i-1, i+2
                while j in range(len(s)) and k in range(len(s)) and s[j] == s[k]:
                    print(j,k)
                    res+=1
                    j-=1
                    k+=1

            # Check palindrome of "1" center
            j,k = i-1, i+1
            while j in range(len(s)) and k in range(len(s)) and s[j] == s[k]:
                    print(j,k)
                    res+=1
                    j-=1
                    k+=1


        return res 
