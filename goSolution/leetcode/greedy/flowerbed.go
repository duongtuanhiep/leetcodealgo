package greedy

/* Algorithm: Greedy algorithm, checks wether can plant at
 * index i, if yes plan and counter++. compare counter >= n
 *
 * Idea: finds maximum amount of flower can be planted in the
 * flowerbed. On every iteration, if it can returns the HIGHEST
 * amount of flower can be planted UP TO THAT POINT then it will
 * be true for the rest of the flowerbed.
 *
 * Can be optimized by increase the index by 1 everytime a flower
 * is planted by the algorithm ssince it is guaranteed that there
 * will be no adjacents, ie if plant in index p then p+1 doesnt have
 * to be checked
**/
func canPlaceFlowers(flowerbed []int, n int) bool {

	// base case, len 1. can plant either 1 or 0 depends on val
	if len(flowerbed) == 1 && flowerbed[0] == 0 || n == 0 {
		return true
	}

	for pos := 0; pos < len(flowerbed); pos++ {
		if flowerbed[pos] == 0 && canPlace(flowerbed, pos) {
			n--
			flowerbed[pos] = 1 // Update this to planted
			pos++
		}
		if n <= 0 {
			return true
		}
	}

	return false
}

func canPlace(arr []int, loc int) bool {

	// Edge case: first elem in arr
	if loc == 0 {
		if arr[loc+1] == 1 {
			return false
		}
		return true
	}

	// last elem
	if loc == len(arr)-1 {
		if arr[loc-1] == 1 {
			return false
		}
		return true
	}

	if arr[loc+1] == 0 && arr[loc-1] == 0 {
		return true
	}
	return false

}
