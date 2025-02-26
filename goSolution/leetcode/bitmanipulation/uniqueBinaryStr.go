package bitmanipulation

import "strconv"

/*

Question 1980: https://leetcode.com/problems/find-unique-binary-string/

Brute force:
- Since the number of possibility is quite small (n <= 16, aka only 2 bytes in size)
we can possibly generate all number representable with 2 bytes and try to see if it
doesn't match any number provided.

Observation:
- Looking at constraint of the question (1 <= n <= 16 and n == nums.length) we can just generate a

**/

func findDifferentBinaryString(nums []string) string {
	presentNum := make(map[int64]struct{})
	for _, val := range nums {
		actualVal, _ := strconv.ParseInt(val, 2, 16)
		presentNum[actualVal] = struct{}{}
	}

	// Get max number representable with bit size len(nums)
	maxNum := (1 << len(nums)) - 1

	for i := 0; i <= maxNum; i++ {
		if _, found := presentNum[int64(i)]; !found {
			resultBase2 := strconv.FormatInt(int64(i), 2)
			for i := len(resultBase2); i < len(nums); i++ {
				resultBase2 = "0" + resultBase2
			}
			return resultBase2
		}
	}

	return ""
}
