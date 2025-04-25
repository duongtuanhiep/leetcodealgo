package dynamicp

/*

Question 39: https://leetcode.com/problems/combination-sum

**/

func combinationSum(candidates []int, target int) [][]int {
	return combSum(candidates, target, 0)
}

func combSum(candidates []int, target, index int) [][]int {
	var res [][]int
	for i := index; i < len(candidates); i++ {
		curNum := candidates[i]
		curArr := []int{}
		for k := 1; k*curNum <= target; k++ {
			curArr = append(curArr, curNum)
			if k*curNum == target {
				res = append(res, curArr)
				break
			}

			if possibles := combSum(candidates, target-curNum*k, i+1); len(possibles) > 0 {
				for _, sub := range possibles {
					sub = append(sub, curArr...)
					res = append(res, sub)
				}
			}
		}
	}

	return res
}
