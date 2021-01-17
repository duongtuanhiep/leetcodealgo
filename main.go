package main

import "fmt"

func main() {
	fmt.Println(countArrangement(2))
}

func countArrangement(n int) int {
	counter, arrStack, curArr := 0, [][]int{}, []int{}
	arrStack = append(arrStack, curArr)
	for len(arrStack) != 0 {
		curArr, arrStack = arrStack[len(arrStack)-1], arrStack[:len(arrStack)-1]

		if len(curArr) == n {
			counter++
			continue
		}

		curArrMap := make(map[int]bool)
		for _, val := range curArr {
			curArrMap[val] = true
		}

		for i := 1; i <= n; i++ {
			if !curArrMap[i] && (i%(len(curArr)+1) == 0 || (len(curArr)+1)%i == 0) {
				newArr := make([]int, len(curArr)+1)
				copy(newArr, curArr)
				newArr[len(newArr)-1] = i
				arrStack = append(arrStack, newArr)
			}
		}
	}
	return counter
}
