package backtracking

/*
Question 17: https://leetcode.com/problems/letter-combinations-of-a-phone-number/

**/

var (
	combo = map[byte][]string{
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
	}
)

func letterCombinations(digits string) []string {
	return combine(digits, 0)
}

func combine(digits string, pos int) []string {
	if pos == len(digits) {
		return nil
	}

	if pos == len(digits)-1 {
		return combo[digits[pos]]
	}

	var res []string
	for _, str := range combine(digits, pos+1) {
		for _, char := range combo[digits[pos]] {
			res = append(res, char+str)
		}
	}
	return res
}
