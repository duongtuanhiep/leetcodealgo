package matrix

// Naive approach
// run max -> then run max v then run max <-- then run max ^. if ^ meet assigned val already then move max ->
func generateMatrix(n int) [][]int {

	res := [][]int{}
	for i := 0; i < n; i++ {
		var subArr []int
		for j := 0; j < n; j++ {
			subArr = append(subArr, 0)
		}
		res = append(res, subArr)
	}

	x, y := 0, 0
	layer := 0
	for num := 1; num <= n*n; num++ {
		if x < len(res)-layer-1 && y == layer || x == y && x == layer {
			res[y][x] = num
			x++
		} else if y < len(res)-layer-1 && x == len(res)-layer-1 {
			res[y][x] = num
			y++
		} else if x <= len(res)-layer-1 && x > layer && y == len(res)-layer-1 {
			res[y][x] = num
			x--
		} else if y <= len(res)-layer-1 && y > layer && x == layer {
			res[y][x] = num
			y--
		}

		if x == layer && y == layer {
			layer++
			x = layer
			y = layer
		}
	}

	return res
}
