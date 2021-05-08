package contest2

// https://leetcode.com/contest/weekly-contest-237/problems/check-if-the-sentence-is-pangram

func checkIfPangram(sentence string) bool {
	char := make(map[rune]bool)
	for _, c := range sentence {
		char[c] = true
	}
	return len(char) == 26
}
