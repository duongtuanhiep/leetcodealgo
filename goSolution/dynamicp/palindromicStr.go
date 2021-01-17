package dynamicp

func longestPalindrome(s string) string {

	charArr := []rune(s)
	if len(charArr) == 0 {
		return ""
	}
	var globalStr string = string(charArr[0])
	var localString string = string(charArr[0])
	var inputStr string
	var isEven = false
	for _, val := range charArr {
		inputStr = inputStr + string(val)
		globalStr, localString = localMaxSubStr(inputStr, globalStr, localString, isEven)
	}
	return globalStr
}

func localMaxSubStr(inputStr string, globalStr string, localStr string, isEven bool) (string, string) {
	charArr := []rune(inputStr)
	lastChar := charArr[len(charArr)-1]

	//base case. Find string like xx
	if len(charArr) >= 2 && lastChar == charArr[len(charArr)-1-1] && !isEven { // 1 element
		localStr = localStr + string(lastChar)
		isEven = true
	} else if isEven {
		if len(charArr)-len(localStr)-2 >= 0 && lastChar == charArr[len(charArr)-len(localStr)-2] {
			localStr = string(lastChar) + localStr + string(lastChar)
		} else if len(charArr)-len(localStr)-1 >= 0 && lastChar == charArr[len(charArr)-len(localStr)-1] {
			localStr = localStr + string(lastChar)
			isEven = false
		} else {
			localStr = string(lastChar)
			isEven = false
		}
	} else if !isEven {
		if len(charArr)-len(localStr)-2 >= 0 && lastChar == charArr[len(charArr)-len(localStr)-2] {
			localStr = string(lastChar) + localStr + string(lastChar)
		} else if len(charArr)-len(localStr)-1 >= 0 && lastChar == charArr[len(charArr)-len(localStr)-1] {
			localStr = localStr + string(lastChar)
			isEven = true
		} else {
			localStr = string(lastChar)
			isEven = false
		}
	} else {
		localStr = string(lastChar)
		isEven = false
	}

	// combine result
	if len(localStr) > len(globalStr) {
		globalStr = localStr
	}

	return globalStr, localStr
}
