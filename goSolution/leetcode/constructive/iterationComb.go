package constructive

import (
	"sort"
)

/*
Question 1286: https://leetcode.com/problems/iterator-for-combination/

**/

type CombinationIterator struct {
	comb []string
	pos  int
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	var charList []rune
	for _, char := range characters {
		charList = append(charList, char)
	}

	sort.SliceStable(charList, func(a, b int) bool { return charList[a] < charList[b] })

	var ci CombinationIterator
	ci.comb = generateStrComb(charList, 0, combinationLength)
	return ci
}

func (this *CombinationIterator) Next() string {
	this.pos++
	return this.comb[this.pos-1]
}

func (this *CombinationIterator) HasNext() bool {
	return this.pos < len(this.comb)
}

/*
Explanation:
pos: next possible position of char on current iteration
length: remaining length to be filled
iterate and call from current position to the last possible position that satify length
because there is no need to say call the last character when there is more than 1 to fill
*/
func generateStrComb(charList []rune, pos, length int) []string {
	if length == 0 || pos >= len(charList) {
		return nil
	}

	var res []string
	for i := pos; i <= len(charList)-length; i++ {
		cur := string(charList[i])
		suffixes := generateStrComb(charList, i+1, length-1)
		if len(suffixes) > 0 {
			for _, sf := range suffixes {
				res = append(res, cur+sf)
			}
		} else {
			res = append(res, cur)
		}
	}

	return res
}

/**
 * Your CombinationIterator object will be instantiated and called as such:
 * obj := Constructor(characters, combinationLength);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
