package graph

func gameOfLife(board [][]int) {
	var neighbors []int
	for x := range board {
		for y := range board[x] {
			neighbors = append(neighbors, neighbor(board, x, y))
		}
	}
	i := 0
	for x := range board {
		for y := range board[x] {
			if board[x][y] == 1 {
				if neighbors[i] < 2 {
					board[x][y] = 0
				}
				if neighbors[i] > 3 {
					board[x][y] = 0
				}
			} else if neighbors[i] == 3 && board[x][y] == 0 {
				board[x][y] = 1
			}
			i++
		}
	}

}

func neighbor(board [][]int, x, y int) int {
	var res int
	lenX := len(board)
	lenY := len(board[x])
	if x < lenX-1 { // bottom
		if board[x+1][y] == 1 {
			res++
		}
	}
	if x < lenX-1 && y < lenY-1 { // bottom right
		if board[x+1][y+1] == 1 {
			res++
		}
	}
	if x < lenX-1 && y > 0 { // bottom left
		if board[x+1][y-1] == 1 {
			res++
		}
	}
	if y < lenY-1 { // right
		if board[x][y+1] == 1 {
			res++
		}
	}
	if y > 0 { // left
		if board[x][y-1] == 1 {
			res++
		}
	}
	if x > 0 && y > 0 { // top left
		if board[x-1][y-1] == 1 {
			res++
		}
	}
	if x > 0 { // top
		if board[x-1][y] == 1 {
			res++
		}
	}
	if x > 0 && y < lenY-1 { // top right
		if board[x-1][y+1] == 1 {
			res++
		}
	}
	return res
}
