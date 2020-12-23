package main

func main() {}
func numIslands(grid [][]byte) int {

	isCheck := make(map[Loc]int)
	queue := []Loc{}
	counter := 0

	for i := range grid {
		for j, val := range grid[i] {
			if val == byte(1) && isCheck[Loc{X: i, Y: j}] == 0 {
				counter++
				queue = append(queue, Loc{X: i, Y: j})
				for len(queue) > 0 {
					u := queue[0].X
					v := queue[0].Y
					isCheck[Loc{X: u, Y: v}] = 1
					if u-1 >= 0 {
						if isCheck[Loc{X: u - 1, Y: v}] == 0 {
							if grid[u-1][v] == byte(1) {
								queue = append(queue, Loc{X: u - 1, Y: v})
							}
						}
					}
					if v-1 >= 0 {
						if isCheck[Loc{X: u, Y: v - 1}] == 0 {
							if (grid[u][v-1]) == byte(1) {
								queue = append(queue, Loc{X: u, Y: v - 1})
							}
						}
					}
					if u+1 <= len(grid)-1 {
						if isCheck[Loc{X: u + 1, Y: v}] == 0 {
							if (grid[u+1][v]) == byte(1) {
								queue = append(queue, Loc{X: u + 1, Y: v})
							}
						}
					}
					if v+1 <= len(grid[u])-1 {
						if isCheck[Loc{X: u, Y: v + 1}] == 0 {
							if (grid[u][v+1]) == byte(1) {
								queue = append(queue, Loc{X: u, Y: v + 1})
							}
						}
					}
					queue = queue[1:]
				}
			}
		}
	}

	return counter
}

type Loc struct {
	X int
	Y int
}
