package wc242

func checkZeroOnes(s string) bool {
	var max1, max0, cur1, cur0 int

	for i := range s {
		if i != 0 {
			if s[i] != s[i-1] {
				if s[i] == '0' {
					cur1 = 0
					cur0 = 1
				} else {
					cur1 = 1
					cur0 = 0
				}
			} else {
				if s[i] == '0' {
					cur0++
				} else {
					cur1++
				}
			}
		} else {
			if s[i] == '0' {
				cur0++
			} else {
				cur1++
			}
		}

		// get max
		if cur1 > max1 {
			max1 = cur1
		}
		if cur0 > max0 {
			max0 = cur0
		}
	}
	return max1 > max0
}
