package divideandconquer

func longestSubstring(s string, k int) int {
	charList := []rune(s)
	return getLongestSub(charList, len(s)-1, 0, k)
}

func getLongestSub(charList []rune, hi, lo, k int) int {

	if hi-lo+1 < k {
		return 0
	}

	countMap := make(map[rune]int)
	for p := lo; p <= hi; p++ {
		countMap[charList[p]]++
	}

	var newMid int
	var midNext int
	for p := lo; p <= hi; p++ {
		if countMap[charList[p]] < k {
			newMid = p
			midNext = p
			for y := p; countMap[charList[y]] < k; y++ {
				midNext++
			}
			break
		} else if p == hi {
			return hi - lo + 1
		}
	}

	return max(getLongestSub(charList, newMid-1, lo, k), getLongestSub(charList, hi, midNext, k))
}

func max(left, right int) int {
	if left > right {
		return left
	}
	return right
}
