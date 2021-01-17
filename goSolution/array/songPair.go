package array

// Brute force solution
// func numPairsDivisibleBy60(time []int) int {
// 	var res int
// 	for i := 0; i < len(time); i++ {
// 		for j := i + 1; j < len(time); j++ {
// 			if (time[i]+time[j])%60 == 0 {
// 				res++
// 			}
// 		}
// 	}
// 	return res
// }

// HINT Solution:
// Calculate, store the reccurence of num%60 in an 60arr length
// Did half a loop to 30.
func numPairsDivisibleBy60(time []int) int {
	reccur := [60]int{}
	var res int
	for _, val := range time {
		reccur[val%60]++
	}
	for i := 0; i <= 30; i++ {
		if i == 0 || i == 30 {
			res += reccur[i] * (reccur[i] - 1) / 2
			continue
		}
		res += reccur[i] * reccur[60-i]
	}
	return res
}
