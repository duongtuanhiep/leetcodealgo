package contest2

// https://leetcode.com/contest/weekly-contest-237/problems/maximum-ice-cream-bars/

func maxIceCream(costs []int, coins int) int {
	var res int
	sort.Ints(costs)
	for _, cost := range costs {
		if coins < cost {
			break
		} else {
			res++
			coins -= cost
		}
	}

	return res
}