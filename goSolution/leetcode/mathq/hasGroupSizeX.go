package mathq

func hasGroupsSizeX(deck []int) bool {

	if len(deck) < 2 {
		return false
	}

	// Count Occ
	occu := make(map[int]int)
	for _, val := range deck {
		occu[val]++
	}

	// Get min, checks if any other groups is divisible for sth thats >=2
	minOcc := occu[deck[0]]
	for _, val := range occu {
		if minOcc > val {
			minOcc = val
		}
	}

	if minOcc == 1 {
		return false
	}

	// get prime divisor
	var divisible []int

	for i := 2; i < minOcc; i++ {
		if minOcc%i == 0 {
			divisible = append(divisible, i)
		}
	}

	for _, div := range divisible {
		i := 0
		for _, val := range occu {
			i++
			if val%div != 0 {
				break
			}
			if i == len(occu)-1 {
				return true
			}
		}
	}

	return false
}
