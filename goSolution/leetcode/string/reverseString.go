package string

/*
Question 917: https://leetcode.com/problems/reverse-only-letters/
**/

func reverseOnlyLetters(s string) string {
	l := len(s) - 1
	cur := []byte(s)
	for f := 0; f < l; {
		if s[l] < 'A' || s[l] > 'z' || s[l] > 'Z' && s[l] < 'a' {
			l--
		} else if s[f] < 'A' || s[f] > 'z' || s[f] > 'Z' && s[f] < 'a' {
			f++
		} else {
			cur[l], cur[f] = cur[f], cur[l]
			f++
			l--
		}
	}
	return string(cur)
}
