package array

import (
	"fmt"
	"math"
	"strconv"
)

/*

Question 3272: https://leetcode.com/problems/find-the-count-of-good-integers/

**/

var digit = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

func countGoodIntegers(n int, k int) int64 {
	pldr := make(map[string]struct{})

	// generate all possible number of len n/2
	halfLen := n / 2
	hasMid := n%2 == 1
	lbound, ubound := int(math.Pow(10, float64(halfLen-1))), int(math.Pow(10, float64(halfLen)))-1
	fmt.Println(lbound, ubound)
	for i := lbound; i <= ubound; i++ {

		// Make final number
		iStr := strconv.Itoa(i)
		riStr := reverse(iStr)

		var possible []string
		if hasMid {
			for _, d := range digit {
				nStr := iStr + d + riStr
				num, _ := strconv.Atoi(nStr)
				if num%k != 0 {
					continue
				}
				possible = append(possible, nStr)
			}
		} else {
			nStr := iStr + riStr
			num, _ := strconv.Atoi(nStr)
			if num%k != 0 {
				continue
			}
			possible = append(possible, nStr)
		}

		for _, combo := range possible {
			for _, valid := range arrange(combo) {
				pldr[valid] = struct{}{}
			}
		}
	}

	return int64(len(pldr))
}

func arrange(str string) []string {
	nCount := make([]int, 10)
	for _, c := range str {
		num, _ := strconv.Atoi(string(c))
		nCount[num]++
	}

	return permutation(nCount, true)
}

func permutation(available []int, first bool) []string {
	var res []string
	for i, cnt := range available {
		if cnt == 0 || i == 0 && first {
			continue
		}

		available[i]--
		combos := permutation(available, false)
		available[i]++

		iStr := strconv.Itoa(i)
		for _, c := range combos {
			res = append(res, iStr+c)
		}
		if len(combos) == 0 {
			res = append(res, iStr) // last character
		}
	}

	return res
}

func reverse(str string) string {
	runeStr := make([]byte, len(str))
	for i := range str {
		runeStr[len(str)-1-i] = str[i]
	}
	return string(runeStr)
}
