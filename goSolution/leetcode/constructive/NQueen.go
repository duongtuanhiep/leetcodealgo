package constructive

/*
Question 51: N-Queens
https://leetcode.com/problems/n-queens/

Idea:
- Create an empty table, try to place first queen at every possible location

Good solution :
https://leetcode.com/submissions/detail/496660885/?from=explore&item_id=3752

**/

// func solveNQueens(n int) [][]string {
// 	var res [][]string
// 	check := make(map[string]bool)

// 	// Try placing "first queen" at every possible pos
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < n; j++ {
// 			remain, table := n, makeTable(n)
// 			remain--
// 			table[i][j] = 'Q'
// 			setRestriction(i, j, table)

// 			// Place if posssible aka blank space.
// 			for x := 0; x < n; x++ {
// 				for y := 0; y < n; y++ {
// 					if table[x][y] == '.' && remain > 0 {
// 						remain--
// 						table[x][y] = 'Q'
// 						setRestriction(x, y, table)
// 					}
// 				}
// 			}

// 			if remain == 0 {
// 				newTable, hash := trim(table)
// 				if !check[hash] {
// 					res = append(res, newTable)
// 					check[hash] = true
// 				}
// 			}
// 		}
// 	}
// 	return res
// }

func solveNQueens(n int) [][]string {
	var res [][]string
	emptyTable := makeTable(n)
	tables := place(0, 0, n, emptyTable)
	for i := range tables {
		res = append(res, trim(tables[i]))
	}
	return res
}

func place(i, j, remain int, curTable [][]rune) [][][]rune {
	if remain == 0 {
		return [][][]rune{curTable}
	}

	var res [][][]rune
	// Place if posssible aka blank space.
	for x := i; x < len(curTable); x++ {
		for y := 0; y < len(curTable); y++ {
			if curTable[x][y] == '.' && remain > 0 { // dfs
				var newTable [][]rune
				for z := range curTable {
					holder := make([]rune, len(curTable))
					copy(holder, curTable[z])
					newTable = append(newTable, holder)
				}
				newTable[x][y] = 'Q'
				setRestriction(x, y, newTable)
				res = append(res, place(x, y, remain-1, newTable)...)
			}
		}
	}

	return res
}

func setRestriction(x, y int, table [][]rune) {
	// up down, left, right
	for i := 0; i < len(table); i++ {
		if i != x {
			table[i][y] = 'x'
		}
		if i != y {
			table[x][i] = 'x'
		}
	}

	// diagonal
	for i := 1; i < len(table); i++ {
		// up left
		if x-i >= 0 && y-i >= 0 {
			table[x-i][y-i] = 'x'
		}
		// up right
		if x-i >= 0 && y+i < len(table) {
			table[x-i][y+i] = 'x'
		}
		// down left
		if x+i < len(table) && y-i >= 0 {
			table[x+i][y-i] = 'x'
		}
		// down right
		if x+i < len(table) && y+i < len(table) {
			table[x+i][y+i] = 'x'
		}
	}
}

func trim(table [][]rune) []string {
	res := []string{}
	for i := range table {
		newLine := ""
		for j := range table[i] {
			if table[i][j] == 'x' {
				newLine += "."
			} else {
				newLine += string(table[i][j])
			}
		}
		res = append(res, newLine)
	}
	return res
}

func makeTable(n int) [][]rune {
	var table [][]rune
	for i := 0; i < n; i++ {
		line := []rune{}
		for j := 0; j < n; j++ {
			line = append(line, '.')
		}
		table = append(table, line)
	}
	return table
}
