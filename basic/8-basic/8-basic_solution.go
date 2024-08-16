package basicEight

import "fmt"

func StringElements(str string) {
	for item := range str {
		fmt.Println(string(str[item]))
	}
}