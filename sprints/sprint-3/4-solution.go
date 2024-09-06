package sprintThreeTaskFourSolution

import "strings"

func detectCapitalUse(word string) bool {

	firstEl := string(word[0])
	subStr := word[1:len(word)]
	if strings.ToUpper(word) == word {
		return true
	} else if strings.ToLower(word) == word {
		return true
	} else if strings.ToUpper(firstEl) == firstEl {
		if strings.ToLower(subStr) == subStr {
			return true
		}
	}

	return false
}