package graph

import (
	"fmt"
	"strconv"
	"strings"
)

// Traditional BFS implemented on 2D grid.
// Store pos in map and queue as "x,y" string
func countBattleships(board [][]byte) int {
	// Visited map
	counter := 0
	visited := make(map[string]bool)
	var queue []string
	var ship byte = 'X'
	for i := range board {
		for j := range board[i] {
			if !visited[toKey(i, j)] {
				if (board[i][j]) == ship {
					counter++
					queue = append(queue, toKey(i, j))
					for len(queue) != 0 {
						i1, j1 := toPos(queue[0])
						visited[toKey(i1, j1)] = true
						if (board[i1][j1]) == ship {
							if i1+1 != len(board)-1 && !visited[toKey(i1+1, j1)] {
								queue = append(queue, toKey(i1+1, j1))
							}
							if j1+1 != len(board[i])-1 && !visited[toKey(i1, j1+1)] {
								queue = append(queue, toKey(i1, j1+1))
							}
						}
						queue = queue[1:]
					}
				}
				visited[toKey(i, j)] = true
			}
		}
	}
	return counter
}

func toPos(this string) (int, int) {
	pos := strings.Split(this, ",")
	a, _ := strconv.Atoi(pos[0])
	b, _ := strconv.Atoi(pos[1])
	return a, b
}

func toKey(x, y int) string {
	return fmt.Sprintf("%v,%v", x, y)
}
