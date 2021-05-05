package bfs

import (
	"fmt"
	"strconv"
	"strings"
)

// func shortestPathBinaryMatrix(grid [][]int) int {

// 	// base/edge
// 	if grid[len(grid)-1][len(grid)-1] == 1 || grid[0][0] == 1 {
// 		return -1
// 	}

// 	// creates queue, steps
// 	var step = 0
// 	var queue []string
// 	visited := grid
// 	// adding "root" to queue
// 	queue = append(queue, toKey(0, 0))

// 	// while discoverable.
// 	for len(queue) != 0 {
// 		length := len(queue) // get current node in layer
// 		step++               // increase step
// 		for n := 0; n < length; n++ {
// 			i, j := toPos(queue[0]) // get current node.
// 			grid[i][j] = step       // assign value path

// 			// adding neighbor to queue,
// 			// only needs to check visited grid for val 0 because visited is the same as grid
// 			if i > 0 && visited[i-1][j] == 0 { // top
// 				visited[i-1][j] = 1
// 				queue = append(queue, toKey(i-1, j))
// 			}
// 			if i > 0 && j > 0 && visited[i-1][j-1] == 0 { // top left
// 				visited[i-1][j-1] = 1
// 				queue = append(queue, toKey(i-1, j-1))
// 			}
// 			if i > 0 && j < len(grid)-1 && visited[i-1][j+1] == 0 { // top right
// 				visited[i-1][j+1] = 1
// 				queue = append(queue, toKey(i-1, j+1))
// 			}
// 			if j > 0 && visited[i][j-1] == 0 { // left
// 				visited[i][j-1] = 1
// 				queue = append(queue, toKey(i, j-1))
// 			}
// 			if j < len(grid)-1 && visited[i][j+1] == 0 { // right
// 				visited[i][j+1] = 1
// 				queue = append(queue, toKey(i, j+1))
// 			}
// 			if i < len(grid)-1 && visited[i+1][j] == 0 { // bottom
// 				visited[i+1][j] = 1
// 				queue = append(queue, toKey(i+1, j))
// 			}
// 			if i < len(grid)-1 && j > 0 && visited[i+1][j-1] == 0 { // bottom left
// 				visited[i+1][j-1] = 1
// 				queue = append(queue, toKey(i+1, j-1))
// 			}
// 			if i < len(grid)-1 && j < len(grid)-1 && visited[i+1][j+1] == 0 { // bottom right
// 				visited[i+1][j+1] = 1
// 				queue = append(queue, toKey(i+1, j+1))
// 			}

// 			// removes from queue
// 			queue = queue[1:]
// 		}
// 	}

// 	// return
// 	if grid[len(grid)-1][len(grid)-1] == 0 {
// 		return -1
// 	}
// 	return grid[len(grid)-1][len(grid)-1]
// }

func toKey(i, j int) string {
	return fmt.Sprintf("%v,%v", i, j)
}

func toPos(pos string) (int, int) {
	posArr := strings.Split(pos, ",")
	i, _ := strconv.Atoi(posArr[0])
	j, _ := strconv.Atoi(posArr[1])
	return i, j
}

func shortestPathBinaryMatrix(grid [][]int) int {

	// base/edge
	if grid[len(grid)-1][len(grid)-1] == 1 || grid[0][0] == 1 {
		return -1
	}

	// creates queue, steps
	var step = 0
	var queue []*Pos
	visited := grid
	// adding "root" to queue
	queue = append(queue, &Pos{0, 0})

	// while discoverable.
	for len(queue) != 0 {
		length := len(queue) // get current node in layer
		step++               // increase step
		for n := 0; n < length; n++ {
			i, j := queue[0].I, queue[0].J // get current node.
			grid[i][j] = step              // assign value path

			// adding neighbor to queue,
			// only needs to check visited grid for val 0 because visited is the same as grid
			if i > 0 && visited[i-1][j] == 0 { // top
				visited[i-1][j] = 1
				queue = append(queue, &Pos{i - 1, j})
			}
			if i > 0 && j > 0 && visited[i-1][j-1] == 0 { // top left
				visited[i-1][j-1] = 1
				queue = append(queue, &Pos{i - 1, j - 1})
			}
			if i > 0 && j < len(grid)-1 && visited[i-1][j+1] == 0 { // top right
				visited[i-1][j+1] = 1
				queue = append(queue, &Pos{i - 1, j + 1})
			}
			if j > 0 && visited[i][j-1] == 0 { // left
				visited[i][j-1] = 1
				queue = append(queue, &Pos{i, j - 1})
			}
			if j < len(grid)-1 && visited[i][j+1] == 0 { // right
				visited[i][j+1] = 1
				queue = append(queue, &Pos{i, j + 1})
			}
			if i < len(grid)-1 && visited[i+1][j] == 0 { // bottom
				visited[i+1][j] = 1
				queue = append(queue, &Pos{i + 1, j})
			}
			if i < len(grid)-1 && j > 0 && visited[i+1][j-1] == 0 { // bottom left
				visited[i+1][j-1] = 1
				queue = append(queue, &Pos{i + 1, j - 1})
			}
			if i < len(grid)-1 && j < len(grid)-1 && visited[i+1][j+1] == 0 { // bottom right
				visited[i+1][j+1] = 1
				queue = append(queue, &Pos{i + 1, j + 1})
			}

			// removes from queue
			queue = queue[1:]
		}
	}

	// return
	if grid[len(grid)-1][len(grid)-1] == 0 {
		return -1
	}
	return grid[len(grid)-1][len(grid)-1]
}

// Revision
type Pos struct {
	I int
	J int
}
