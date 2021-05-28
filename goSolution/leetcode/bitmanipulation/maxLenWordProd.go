package bitmanipulation

/*
Question 318: https://leetcode.com/problems/maximum-product-of-word-lengths/


Idea:
- We will have to check every possible "combination" of 2 words. If it is valid
then we can calculate the multiplication.
- To shorten the amount of time it takes to compare each 2 words, instead of comparing
by each character of the string, we can store the mask of them and with AND operation we will
be able to compare it in constant time.

Consider a 26-bit integer, we can store each character in a string in the bit representing
it placement in the alphabet (1 equal to exist, 0 equal to not exist)
Example:
- A string of just "a" can be represented as 1
	(character a is at first place in the alphabet)

- A string of "ac" can be represented as 101 in binary = 5 in decimal
	- character a is at first place in alphabet -> first bit is 1
	- character b does not exist -> second bit is 0
	- character c exist -> third bit is 0

**/

// Hash map solution
func maxProductHashMap(words []string) int {
	var chars []map[rune]bool
	for _, word := range words {
		cmap := make(map[rune]bool)
		for _, c := range word {
			cmap[c] = true
		}
		chars = append(chars, cmap)
	}
	var res, cur int
	for i := range chars {
		for j := range chars {
			if i == j {
				continue
			}
			valid := true
			for character := range chars[i] {
				if chars[i][character] == chars[j][character] {
					valid = false
					break
				}
			}
			if valid {
				cur = len(words[i]) * len(words[j])
				if cur > res {
					res = cur
				}
			}
		}
	}

	return res
}

// Hash map solution
func maxProduct(words []string) int {
	var masks []int
	for _, word := range words {
		mask := 0
		for _, char := range word {
			mask = mask | 1<<int(char-'a')
		}
		masks = append(masks, mask)
	}

	var res int
	for i := range masks {
		for j := range masks {
			if i == j {
				continue
			}
			if masks[i]&masks[j] == 0 {
				res = max(res, len(words[i])*len(words[j]))
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
