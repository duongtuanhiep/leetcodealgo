package stack

/*
Question 1209: https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string-ii/

Solution:
- Simple:
	scan through the whole string, removes anything that is possible to removes.
	Return when there is nothing to removes.
	Runtime O(N^K^2)
- Optimized:

**/

func removeDuplicates(s string, k int) string {
	// var newRes string
	var cur string
	for {
		for i := 0; i < len(s); i++ {
			if isDuplicate(i, i+k, s) {
				i = i + k - 1
			} else {
				cur = cur + string(s[i])
			}
		}
		// fmt.Println(cur)
		if s == cur {
			break
		}
		s, cur = cur, ""
	}
	return s
}

func isDuplicate(start, end int, s string) bool {
	if start >= len(s) || end > len(s) {
		return false
	}
	last := s[start]
	for i := start; i < end; i++ {
		if s[i] != last {
			return false
		}
	}
	return true
}
