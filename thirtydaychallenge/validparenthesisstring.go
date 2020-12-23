package thirtydaychallenge

func checkValidString(s string) bool {

	// ( = 1
	// ) = -1
	// * = + 1
	if len(s) == 0 {
		return true
	}
	var low, high = 0, 0
	for _, val := range s {
		if string(val) == "(" {
			low++
		} else {
			low--
		}
		if string(val) != ")" {
			high++
		} else {
			high--
		}
		if high < 0 {
			return false
		}
		if low < 0 {
			low = 0
		}
	}
	if low == 0 {
		return true
	}
	return false
}

// ((*)(*)((*
// (*))
// ((*)(*))((*
// ((*)(*)))(((*
