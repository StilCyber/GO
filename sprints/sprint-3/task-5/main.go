package main

import "strings"

func lengthOfLastWord(s string) int {

	stringWithoutEndSpaces := strings.Trim(s, " ")

	arrayString := strings.Split(stringWithoutEndSpaces, " ")
	lastWord := arrayString[len(arrayString)-1]

	return len(lastWord)
}