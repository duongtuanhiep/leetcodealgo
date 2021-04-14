package array

/*
Question 1704: https://leetcode.com/problems/determine-if-string-halves-are-alike/


**/

var (
	allowed = map[byte]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}
)

func halvesAreAlike(s string) bool {
	left, right := 0, 0
	for i, j := 0, len(s)/2; i < len(s)/2; i, j = i+1, j+1 {
		if allowed[s[i]] {
			left++
		}
		if allowed[s[j]] {
			right++
		}
	}
	return left == right
}
