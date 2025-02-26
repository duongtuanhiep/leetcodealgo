package stupid

/*
Question 2689: https://leetcode.com/problems/find-the-punishment-number-of-an-integer/

Approach: Brute Force
Try to compute every single number from 1 to n, and check if the square of the number is equal to the sum of the digits of the number.

Approach: DP & Backtracking
Few considerations:
- Instead of trying to compute all perms, try to actively compute sums of digits for each perms and drops when it obviously not match
- Instead of using string, use int with `mod` and `div` to extract the string.
- Instead of doing top down (compute left most first), do bottom up (compute right most first)

// Brute force solution

	func punishmentNumber(n int) int {
		var res int

		for i := 1; i <= n; i++ {
			sqrt := i * i
			sqrtStr := strconv.Itoa(sqrt)
			perms := generatePermutation(sqrtStr)

			for _, perm := range perms {
				if perm == i {
					fmt.Println(i, sqrt)
					res += sqrt
					break
				}
			}
		}

		return res
	}

	func generatePermutation(inStr string) []int {
		if len(inStr) == 0 {
			return []int{}
		}

		var res []int

		for i := 0; i < len(inStr); i++ {
			curStr := inStr[0 : i+1]
			curInt, _ := strconv.Atoi(curStr)

			remainPerms := generatePermutation(inStr[i+1:])
			if len(remainPerms) == 0 {
				res = append(res, curInt)
			} else {
				for _, perm := range remainPerms {
					res = append(res, curInt+perm)
				}
			}

		}

		return res
	}
*/

// Backtracking
func punishmentNumber(n int) int {
	var res int

	for i := 1; i <= n; i++ {
		if isPartitionValid(i*i, i) {
			res += i * i
		}
	}

	return res
}

func isPartitionValid(num, target int) bool {
	if target < 0 || num < target {
		return false
	}

	if num == target {
		return true
	}

	return isPartitionValid(num/10, target-num%10) || isPartitionValid(num/100, target-num%100) || isPartitionValid(num/1000, target-num%1000)
}
