package dfs

import (
	"bytes"
)

/*
Question 1079: https://leetcode.com/problems/letter-tile-possibilities/description/

Observation:
- If we just try our best to generate all cases
**/

func numTilePossibilities(tiles string) int {
	combos := make(map[string]bool)
	_ = generate(tiles, combos)

	return len(combos)
}

func generate(tiles string, combos map[string]bool) []string {
	var curStr string
	var res []string
	for i, r := range tiles {
		curStr = string(r)
		if !combos[curStr] {
			combos[curStr] = true
		}
		res = append(res, curStr)
		var buf bytes.Buffer
		buf.Grow(len(tiles) - 1)
		buf.WriteString(tiles[:i])
		buf.WriteString(tiles[i+1:])
		nexts := generate(buf.String(), combos)
		for _, next := range nexts {
			if !combos[next] {
				combos[next] = true
			}
			res = append(res, next)
			for i := 0; i <= len(next); i++ {
				var sb bytes.Buffer
				sb.Grow(len(next) + len(curStr)) // Preallocate required space
				sb.WriteString(next[:i])
				sb.WriteString(curStr)
				sb.WriteString(next[i:])
				final := next[:i] + curStr + next[i:]
				if !combos[final] {
					combos[final] = true
					res = append(res, final)
				}
			}
		}
	}
	return res
}
