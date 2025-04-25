package slidingwindows

/*
Question 3: https://leetcode.com/problems/longest-substring-without-repeating-characters/


*/

/*
Idea: optimmized brute force
Exaustive search: Do linear scan through the array, have a map to count reccurence. On
each element of outer loop does a reccurence count. Stop if any elements has > 1 occurence.
Runtime O(N^2), Space O(N)
*
*/
func lengthOfLongestSubstring(s string) int {
	var maxStr int
	for i := 0; i < len(s); i++ {
		repeat, curMax := make(map[byte]bool), 0
		for j := i; j < len(s) && !repeat[s[j]]; j++ {
			curMax++
			repeat[s[j]] = true
		}
		if curMax > maxStr {
			maxStr, curMax = curMax, 0
		}
	}
	return maxStr
}

/*
Observation:
- Any 2 duplicate elements can not belong to the same substring. Say for character
'a' that exist at index 0 and 4 will certainly belong to different longest substring
that has element 'a'
- for every newly checked element i+1 we can decide at that moment what could be the
the longest substring that i+1 a part of from 0 to i+1
- can creates another array to store the index of the last appearance of that character.
- Keep track of the start index of the current possible subString that the procesing
element belongs to.

This is my own solution that is somewhere in between of approach 2 and 3 from the solution
in terms of space complexity.
https://leetcode.com/problems/longest-substring-without-repeating-characters/solution/

Variation of sliding windows. The goal is to compute the longest sub-string that the
current elements in the interation belongs to. returns max

Runtime O(N), Space O(N)

TODO: Optimize by instead of using an alternative array to store "last" occurence of
character we can create a hash map and do this on the fly.
*
*/
func lengthOfLongestSubstring(s string) int {
	var lcp []int // Array that store index of last reccurence of a char
	newestRec := make(map[byte]int)

	// Preproccessing
	for i := range s {
		if _, found := newestRec[s[i]]; !found {
			lcp = append(lcp, -1)
		} else {
			lcp = append(lcp, newestRec[s[i]])
		}
		newestRec[s[i]] = i
	}

	var maxStr, localStr, start = 0, 0, 0
	for i := 0; i < len(s); i++ {
		if lcp[i] < start {
			localStr = localStr + 1
		} else {
			localStr = localStr - (lcp[i] - start)
			start = lcp[i] + 1
		}
		if localStr > maxStr {
			maxStr = localStr
		}
	}

	return maxStr
}
