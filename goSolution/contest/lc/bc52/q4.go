package bc52

/*
[7,7,7,7,7,7,7]

**/
func sumOfFlooredPairs(nums []int) int {
	var res int

	count := make([]int, 100001)
	for _, val := range nums {
		count[val]++
	}

	for i := len(count) - 1; i > 0; i-- {
		if count[i] != 0 {
			res += count[i] * count[i]
			for j := i - 1; j > 0; j-- {
				if count[j] != 0 {
					res += (count[i] * count[j]) * (i / j)
				}
			}
		}
	}

	return res % (1000000000 + 7)
}
