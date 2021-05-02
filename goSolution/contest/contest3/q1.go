package contest3

import "strconv"

var (
	charmap = map[rune]int{
		'a': 0, 'b': 1, 'c': 2, 'd': 3, 'e': 4, 'f': 5, 'g': 6, 'h': 7, 'i': 8, 'j': 9, 'k': 10, 'l': 11,
		'm': 12, 'n': 13, 'o': 14, 'p': 15, 'q': 16, 'r': 17, 's': 18, 't': 19, 'u': 20, 'v': 21, 'w': 22, 'x': 23, 'y': 24, 'z': 25}
	charPos = []rune{
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
)

func replaceDigits(s string) string {

	newStr := ""
	for i, char := range s {
		if i%2 == 0 {
			newStr = newStr + string(char)
		} else {
			num, _ := strconv.Atoi(string(char))
			newStr = newStr + string(shift(rune(s[i-1]), num))
		}
	}
	return newStr
}

func shift(s rune, i int) rune {
	return charPos[charmap[s]+i]
}
