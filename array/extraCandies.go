package array

func kidsWithCandies(candies []int, extraCandies int) []bool {
	// 2 pass
	maxCandy := 0
	for i := range candies {
		if candies[i] > maxCandy {
			maxCandy = candies[i]
		}
	}
	var res []bool
	for _, val := range candies {
		if val+extraCandies >= maxCandy {
			res = append(res, true)
		} else {
			res = append(res, false)
		}
	}
	return res
}
