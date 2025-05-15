"""

Question 1358: https://leetcode.com/problems/number-of-substrings-containing-all-three-characters/

Solution: 
Consider : l,r denotes the start,end of the current str. 
- if "___" is a valid string, we know that all the "___X","___XX", "___XXX" is going to be valid too.
- that means that the total number of valid sub str starting at index i is going to be end - j 
- we can extends l until theres a valid string, and then 

Some optimization: 
- Instead of a dict coulve just use the arrays and use `ord(x) - ord(a)` to get the pos.

"""


class Solution:
    def numberOfSubstrings(self, s: str) -> int:
        count = dict()
        i, j, res = 0, 0, 0
        count[s[j]] = count.get(s[j], 0) + 1

        while 0 <= i <= j < len(s):
            if not self.valid(count):
                j += 1
                if j < len(s):
                    count[s[j]] = count.get(s[j], 0) + 1
            else:
                res += len(s) - j
                count[s[i]] = count.get(s[i], 0) - 1
                i += 1

        return res

    def valid(self, count: dict) -> bool:
        if count.get("a", 0) and count.get("b", 0) and count.get("c", 0):
            return True
        return False


if __name__ == "__main__":
    cases = ["abcabc", "aaacb", "abc", "abcccccc", "aabccc"]

    for c in cases:
        print(Solution().numberOfSubstrings(c))
