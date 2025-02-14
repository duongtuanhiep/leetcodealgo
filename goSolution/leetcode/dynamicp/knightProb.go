package dynamicp

/**
Question 688: https://leetcode.com/problems/knight-probability-in-chessboard/

Precalculate the board so that its quicker to calculate later on.
*/

// func knightProbability(n int, k int, row int, column int) float64 {
// 	// Pre calculate prob-table
// 	board := make([][]float64, n)
// 	for i := range board {
// 		board[i] = make([]float64, n)
// 	}

// 	for i := 0; i < n; i++ {
// 		for j := 0; j < n; j++ {
// 			var positiveCase int
// 			if i+1 < n && j+2 < n {
// 				positiveCase++
// 			}
// 			if i+1 < n && j-2 >= 0 {
// 				positiveCase++
// 			}
// 			if i+2 < n && j+1 < n {
// 				positiveCase++
// 			}
// 			if i+2 < n && j-1 >= 0 {
// 				positiveCase++
// 			}
// 			if i-1 >= 0 && j+2 < n {
// 				positiveCase++
// 			}
// 			if i-1 >= 0 && j-2 >= 0 {
// 				positiveCase++
// 			}
// 			if i-2 >= 0 && j+1 < n {
// 				positiveCase++
// 			}
// 			if i-2 >= 0 && j-1 >= 0 {
// 				positiveCase++
// 			}
// 			board[i][j] = float64(positiveCase) / 8.0
// 		}
// 	}

// 	return kProb(row, column, k, board)
// }

// func kProb(posX, posY, k int, board [][]float64) float64 {
// 	if posX < 0 || posY < 0 || posX >= len(board) || posY >= len(board) {
// 		return 0.0
// 	} else if k == 0 {
// 		return 1.0
// 	} else if k == 1 {
// 		return board[posX][posY]
// 	}
// 	return kProb(posX-2, posY-1, k-1, board)/8 +
// 		kProb(posX-2, posY+1, k-1, board)/8 +
// 		kProb(posX+2, posY-1, k-1, board)/8 +
// 		kProb(posX+2, posY+1, k-1, board)/8 +
// 		kProb(posX-1, posY-2, k-1, board)/8 +
// 		kProb(posX+1, posY-2, k-1, board)/8 +
// 		kProb(posX-1, posY+2, k-1, board)/8 +
// 		kProb(posX+1, posY+2, k-1, board)/8
// }

func knightProbability(n int, k int, row int, column int) float64 {
	// Pre calculate prob-table
	lastBoard, currentBoard := make([][]float64, n), make([][]float64, n)
	for i := range lastBoard {
		lastBoard[i], currentBoard[i] = make([]float64, n), make([]float64, n)

		// pre-fill last board
		for j := range lastBoard[i] {
			lastBoard[i][j] = 1.0
		}
	}

	for count := 0; count < k; count++ {
		// calculate current iteration based from last iteration
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				var prob float64
				if i+1 < n && j+2 < n {
					prob += lastBoard[i+1][j+2] / 8
				}
				if i+1 < n && j-2 >= 0 {
					prob += lastBoard[i+1][j-2] / 8
				}
				if i+2 < n && j+1 < n {
					prob += lastBoard[i+2][j+1] / 8
				}
				if i+2 < n && j-1 >= 0 {
					prob += lastBoard[i+2][j-1] / 8
				}
				if i-1 >= 0 && j+2 < n {
					prob += lastBoard[i-1][j+2] / 8
				}
				if i-1 >= 0 && j-2 >= 0 {
					prob += lastBoard[i-1][j-2] / 8
				}
				if i-2 >= 0 && j+1 < n {
					prob += lastBoard[i-2][j+1] / 8
				}
				if i-2 >= 0 && j-1 >= 0 {
					prob += lastBoard[i-2][j-1] / 8
				}
				// fmt.Println(prob)
				currentBoard[i][j] = prob
			}
		}

		// fmt.Println(lastBoard, currentBoard)

		// Copy current iteration board to last iteration
		for i := range currentBoard {
			copy(lastBoard[i], currentBoard[i])
		}
	}

	return lastBoard[row][column]
}
