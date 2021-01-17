package greedy

/*
Problem 392: https://leetcode.com/problems/is-subsequence/

Statement: check if string t is a subsequence of t (by deleting
some elements but not changing the order of any elements)

Idea: Not changing the order of the character but possible to
"skip" characters in between suggesting that we can have an
iterative approach, scanning through the string with counter i,
increase j when there is a matching rune in string, drop when
it is not matched.
At the end, return if our counter j has same length as strings
t.

Greedy algorithm idea:
On any string s and string t, if t[0:a] is a subsequence of
s[0:c], t[a:b] is subsequence of s[c:d] then t[0:b] is subsequence
of s[0:d]. This should also be the case for the rest of the string.
This allows a greedy approach: checks for the matching iteratively,
if it matchs then its will be included in the end result.
(Needs further explanation - ?)


UPDATE:
- No need to convert to rune. It is rune by default if you index
the string
**/

func isSubsequence(s string, t string) bool {
	var j, i int
	for i, j = 0, 0; i < len(t) && j < len(s); i++ {
		if t[i] == s[j] {
			j++
		}
	}
	return j == len(s)
}
