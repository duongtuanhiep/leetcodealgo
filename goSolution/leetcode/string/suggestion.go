package string

import (
	"sort"
	"strings"
)

/*
Question 1268: https://leetcode.com/problems/search-suggestions-system/


Idea:
- Sort + BS: sort the whole array, find the leftmost elem that has the same fixes
as other. Dont do the stupid linear scan :-( (me)
- Trie

https://leetcode.com/problems/search-suggestions-system/solution/
**/

func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)
	var res [][]string
	for i := range searchWord {
		var suggest []string
		for j := bs(products, searchWord[:i+1]); j < len(products) && len(suggest) < 3; j++ {
			if strings.Index(products[j], searchWord[:i+1]) == 0 {
				suggest = append(suggest, products[j])
			}
		}
		res = append(res, suggest)
	}
	return res
}

func bs(prod []string, fixes string) int {
	hi, lo := len(prod)-1, 0
	for lo < hi {
		mid := (hi + lo) / 2
		if prod[mid] < fixes {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}
