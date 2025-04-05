package dynamicp

/*

Question 2140: https://leetcode.com/problems/solving-questions-with-brainpower/?envType=daily-question&envId=2025-04-04

Intuition: For each [point, brainpower] index i, we have formula
result[i] = max(index[i] (if cant solve anymore), index[i+1], index[i] + index[i+brainpower])

More on problem idea: https://en.wikipedia.org/wiki/Knapsack_problem
**/

func mostPoints(questions [][]int) int64 {
	maxPoint := make([]int, len(questions))
	for i := len(questions) - 1; i >= 0; i-- {
		if i == len(questions)-1 {
			maxPoint[i] = questions[i][0]
		} else {
			maxPoint[i] = max(maxPoint[i+1], questions[i][0])
			nextSolvableIndex := i + questions[i][1] + 1
			if nextSolvableIndex < len(questions) {
				maxPoint[i] = max(maxPoint[i], questions[i][0]+maxPoint[nextSolvableIndex])
			}
		}
	}

	return int64(maxPoint[0])
}
