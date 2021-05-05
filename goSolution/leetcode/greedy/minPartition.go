package greedy

/*
Question 1689: https://leetcode.com/problems/partitioning-into-minimum-number-of-deci-binary-numbers/

Idea:
	- The question requires you to reconstruct a number in SUM of deci-binary number, we notice that
	any specific integer character n (0 -> 9) can only be computed by deci-binary of value of such as
	n = 1 * n. Meaning that the possible minimum solution is bounded by the largest character, if there
	exist number 9 then no matter how many deci-binary it takes to represent other number you will need
	at least 9 number.

Optimization:
	- if there is a 9 we can return 9 right away without checking the rest
	- Instead of converting from rune (golang's Char) to integer, we can use its integer representation
	right away

Note: Golang Rune encoding scheme:
	- rune('0') = 48
	- rune('9') = 57
**/

func minPartitions(n string) int {
	var res int
	for i := range n {
		if int(n[i]) > res {
			res = int(n[i])
		}
		if res == 57 { // encoding of character 9
			return 9
		}
	}
	return res - 48
}
