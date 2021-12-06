package cpSetup

// Left most 
func bs(prod []string, fixes string) int {
	hi, lo := len(prod)-1, 0
	for lo < hi {
		mid := (hi + lo) / 2
		if prod[mid] < fixes {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}
