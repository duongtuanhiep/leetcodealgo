package backtracking

/*
Question 526: https://leetcode.com/problems/beautiful-arrangement/

Idea: DFS/Backtracking using stack
**/

// DFS with Stack on slices of slices
// func countArrangement(n int) int {
// 	counter, arrStack, curArr := 0, [][]int{}, []int{}
// 	arrStack = append(arrStack, curArr)
// 	for len(arrStack) != 0 {
// 		curArr, arrStack = arrStack[len(arrStack)-1], arrStack[:len(arrStack)-1]

// 		if len(curArr) == n {
// 			counter++
// 			continue
// 		}

// 		curArrMap := make(map[int]bool)
// 		for _, val := range curArr {
// 			curArrMap[val] = true
// 		}

// 		for i := 1; i <= n; i++ {
// 			if !curArrMap[i] && (i%(len(curArr)+1) == 0 || (len(curArr)+1)%i == 0) {
// 				newArr := make([]int, len(curArr)+1)
// 				copy(newArr, curArr)
// 				newArr[len(newArr)-1] = i
// 				arrStack = append(arrStack, newArr)
// 			}
// 		}
// 	}
// 	return counter
// }

// a more elegant way with recursion and no array created. Translated from solution.
func countArrangement(n int) int {
	visited := make([]bool, n+2)
	return count(n, 1, 0, visited)
}

func count(n, pos, counter int, visited []bool) int {
	if pos > n {
		return counter + 1
	}
	for i := 1; i <= n; i++ {
		if !visited[i] && (pos%i == 0 || i%pos == 0) {
			visited[i] = true
			counter = count(n, pos+1, counter, visited)
			visited[i] = false
		}
	}
	return counter
}
