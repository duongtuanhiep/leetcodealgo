package hashmap

import "strconv"

/*
Question 49: https://leetcode.com/problems/group-anagrams

*
*/

func groupAnagrams(strs []string) [][]string {

	strCache := make(map[string][]string)
	for i := range strs {
		key := toKey(strs[i])
		if strCache[key] == nil {
			strCache[key] = []string{strs[i]}
		} else {
			strCache[key] = append(strCache[key], strs[i])
		}
	}

	res := [][]string{}
	for _, v := range strCache {
		res = append(res, v)
	}
	return res
}

func toKey(str string) string {
	runeArr := make([]int, 26)
	for _, c := range str {
		runeArr[c-'a']++
	}

	var res string
	for i := range runeArr {
		if runeArr[i] != 0 {
			res += strconv.Itoa(i) + ":" + strconv.Itoa(runeArr[i]) + ";"
		}
	}

	return res
}
