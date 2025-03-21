package sprintThreeTaskThreeSolution

import "strings"

func mostWordsFound(sentences []string) int {
    result := 0

    for item := range sentences {
        num := len(strings.Split(sentences[item], " "))
        if(num > result) {
            result = num
        }
    }

    return result
}