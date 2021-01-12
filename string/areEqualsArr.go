package string

import "strings"

/*
Question 1662: https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent/submissions/
self explantory
**/

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	return strings.Join(word1, "") == strings.Join(word2, "")
}
