package greedy

/*
Question 1405: https://leetcode.com/problems/longest-happy-string/

Idea: Maintaining a heap of current available characters available
to build the longest strings.
Each time remove max value, add character max once, then add in max--
if last 2 is max character removes max, removes new max, add in
second max--, first old max
Stop when len(heap) == 1 && last2 == last char in heap
Runtime: O(a + b + c) on case when can fill all the available a,b,c.
Number of heap oparetaion is hard bounded since there are predicted number
of elements (only a,b,c).

Building a heap and utilize it is highly reccomended BUT...

Using Go standard library is slow (dealing with type interface and
type asertion of type) and requires more work setting up required functions.

A "dumber" way to solve is to instead of mantaining a heap (more
generic solution), since we know there will be at most 3 types of element
'a','b' and 'c' might as well do it manually. This, despite lots of
additional tracking is still faster than the heap implementation (this gets
100% runtime and 100% space)
**/

func longestDiverseString(a int, b int, c int) string {
	var curStr []rune
	counter := make(map[rune]int)
	counter['a'] = a
	counter['b'] = b
	counter['c'] = c

	for hasLeft(counter) {
		next := firstMax(counter)
		if len(curStr) < 2 || (curStr[len(curStr)-2] != next || curStr[len(curStr)-1] != next) {
			curStr = append(curStr, next)
			update(counter, next)
		} else {
			secondNext := secondMax(counter)
			if counter[secondNext] == 0 {
				break
			}
			curStr = append(curStr, secondNext)
			update(counter, secondNext)
		}
	}

	return string(curStr)
}

func hasLeft(c map[rune]int) bool {
	var res = false
	for _, count := range c {
		if count > 0 {
			res = true
			break
		}
	}
	return res
}

func firstMax(c map[rune]int) rune {
	if c['a'] >= c['b'] && c['a'] >= c['c'] {
		return 'a'
	} else if c['b'] >= c['a'] && c['b'] >= c['c'] {
		return 'b'
	}
	return 'c'
}

func secondMax(c map[rune]int) rune {
	if c['a'] >= c['b'] && c['a'] >= c['c'] {
		if c['b'] >= c['c'] {
			return 'b'
		}
		return 'c'
	} else if c['b'] >= c['a'] && c['b'] >= c['c'] {
		if c['a'] >= c['c'] {
			return 'a'
		}
		return 'c'
	}
	if c['a'] >= c['b'] {
		return 'a'
	}
	return 'b'
}

func update(counter map[rune]int, char rune) {
	counter[char]--
}
