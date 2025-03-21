package main
import "strings"

func findWords(words []string) []string {
    var result []string

    firstRow := "qwertyuiop"
    secondRow := "asdfghjkl"
    thirdRow := "zxcvbnm"

    for _, item := range words {
        itemArr := strings.Split(strings.ToLower(string(item)), "")

        flag := true

        if(strings.Contains(firstRow, itemArr[0])) {
            for _, letter := range itemArr {
                if(!strings.Contains(firstRow, letter)) {
                    flag = false
                    break
                }
            }

        } else if (strings.Contains(secondRow, itemArr[0])) {
            for _, letter := range itemArr {
                if(!strings.Contains(secondRow, letter)) {
                    flag = false
                    break
                }
            }

        } else if (strings.Contains(thirdRow, itemArr[0])) {

            for _, letter := range itemArr {
                if(!strings.Contains(thirdRow, letter)) {
                    flag = false
                    break
                }
            }
            
        }

        if(flag) {
            result = append(result, item)
        }
    }

    return result
}