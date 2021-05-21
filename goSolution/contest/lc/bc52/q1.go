package bc52

import "strconv"

func sortSentence(s string) string {
	s = s + string(' ')
	strMap := make([]string, 10)

	last := 0
	for i := range s {
		if s[i] == ' ' {
			curStr := s[last : i-1]
			num, _ := strconv.Atoi(string(s[i-1]))
			strMap[num] = curStr
			last = i + 1
		}
	}

	res := ""
	for i := 1; i < len(strMap) && strMap[i] != ""; i++ {
		res += strMap[i] + " "
	}

	return res[:len(res)-1]
}
