
"""

Question 3306:

Observation: 
- Basically asked to find substring contains [a,e,i,o,u] and exactly k wildcard
- Can we just perform pattern matching, then we can see 

Algo (Less smart, more space but O(N))
- Define nextC denoting the next consonants we can find from any position
- Define l,r =0,0
- While current s[l:r] not valid, keep extending r until either consonent exceed k, after which start srhinking l
- When s[l:r] valid, shrink l and adds nextConsonant[r] - r to the sum 
(nextConsonant[r] should point to index of next c, r is current rightbound, since we can have that many combinations that is still valid)
- After s[l:r] not valid anymore, start extending r again. 
"""


class Solution:
    def _isVowel(self, c: str) -> bool:
        return c in ["a", "e", "i", "o", "u"]

    def _atLeastK(self, word: str, k: int) -> int:
        num_valid_substrings = 0
        start = 0
        end = 0
        # keep track of counts of vowels and consonants
        vowel_count = {}
        consonant_count = 0

        # start sliding window
        while end < len(word):
            # insert new letter
            new_letter = word[end]

            # update counts
            if self._isVowel(new_letter):
                vowel_count[new_letter] = vowel_count.get(new_letter, 0) + 1
            else:
                consonant_count += 1

            # shrink window while we have a valid substring
            while len(vowel_count) == 5 and consonant_count >= k:
                num_valid_substrings += len(word) - end
                start_letter = word[start]
                if self._isVowel(start_letter):
                    vowel_count[start_letter] = (
                        vowel_count.get(start_letter) - 1
                    )
                    if vowel_count.get(start_letter) == 0:
                        vowel_count.pop(start_letter)
                else:
                    consonant_count -= 1
                start += 1

            end += 1

        return num_valid_substrings

    def countOfSubstrings(self, word: str, k: int) -> int:
        return self._atLeastK(word, k) - self._atLeastK(word, k + 1)
