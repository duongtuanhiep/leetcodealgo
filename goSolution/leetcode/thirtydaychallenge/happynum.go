package thirtydaychallenge

import (
	"strconv"
)

func isHappy(n int) bool {
	var fast, slow = n, n
	if f(n) == 1 {
		return true
	}
	for true {
		slow = f(slow)
		fast = f(f(fast))
		if slow == fast {
			return false
		}
		if slow == 1 || fast == 1 {
			return true
		}
	}
	return true
}

func f(n int) int {
	var result = 0
	var inputStr = strconv.Itoa(n)
	for _, val := range inputStr {
		intVal, _ := strconv.Atoi(string(val))
		result += intVal * intVal
	}
	return result
}
