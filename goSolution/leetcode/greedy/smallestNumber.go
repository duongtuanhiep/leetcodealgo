package greedy

import "strconv"

/*

Question 2375: https://leetcode.com/problems/construct-smallest-number-from-di-string/

Observation:
- We know the length & maximum number from the length

- Try to put as small number as possible, put them in stack for each decrease

**/

func smallestNumber(pattern string) string {
	curChar := make([]int, len(pattern)+1)
	curChar[0] = 1
	var accendingStack []int

	curNum := 1
	for i, r := range pattern {
		curNum++
		if r == 'D' { // We just put it into a stack and handle later
			accendingStack = append(accendingStack, curNum)
		} else if r == 'I' { // We pop the stack to "compensate" for earlier
			if len(accendingStack) > 0 {
				curChar[i] = curChar[i-len(accendingStack)] // Swap the last 'I' value to the last visited index
				for len(accendingStack) > 0 {
					curChar[i-len(accendingStack)] = accendingStack[len(accendingStack)-1]
					accendingStack = accendingStack[:len(accendingStack)-1]
				}
			}
			curChar[i+1] = curNum
		}

	}

	// If last pattern was 'D', pop the stack manually
	if len(accendingStack) > 0 {
		curChar[len(pattern)] = curChar[len(pattern)-len(accendingStack)] // Value at last 'I'
		for len(accendingStack) > 0 {
			curChar[len(pattern)-len(accendingStack)] = accendingStack[len(accendingStack)-1]
			accendingStack = accendingStack[:len(accendingStack)-1]
		}

	}

	// make a string
	var res string
	for _, val := range curChar {
		res += strconv.Itoa(val)
	}

	return res
}
