package dynamicp

/*
Question 1641: https://leetcode.com/problems/count-sorted-vowel-strings/


Available string arr [a, e, i, o, u]
Can think of making a number 'abc' so that a <= b <= c comprised of 1,2,3,4,5
**/

/*
Idea: Instead of creating string we can denote vowel as int (1->5 ) and recursively
"place" possible solution.
*/
func countVowelStrings(n int) int {
	return canPlace(1, 0, n)
}

func canPlace(last, curLen, n int) int {
	var result int
	if curLen == n {
		return 1
	}
	for i := last; i <= 5; i++ {
		result += canPlace(i, curLen+1, n)
	}
	return result
}
