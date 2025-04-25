package slidingwindows

/*
Question 696: https://leetcode.com/problems/count-binary-substrings/

Idea: Keep count of the number of 0 and 1 in sliding windows

*/

/*
func countBinarySubstrings(s string) int {
	var res int
	counter := make(map[rune]int)
	var last rune
	for _, char := range s {
		if counter[char] > 0 && char != last {
			res += minVal(counter[char], counter[last])
			counter[char] = 0
		}
		counter[char]++
		last = char
	}

	return res + minVal(counter['0'], counter['1'])
}
*/

func countBinarySubstrings(s string) int {
	var res, cur0, cur1 int
	var last rune
	for _, char := range s {
		if char == '0' {
			if cur1 != 0 && last != char {
				res += minVal(cur0, cur1)
				cur0 = 0
			}
			cur0++
		} else if char == '1' {
			if cur0 != 0 && last != char {
				res += minVal(cur0, cur1)
				cur1 = 0
			}
			cur1++
		}
		last = char
	}

	return res + minVal(cur0, cur1)
}

func minVal(left, right int) int {
	if left > right {
		return right
	}
	return left
}
