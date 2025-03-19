package main

import "fmt"

func main() {
	a := 5000
	for i := 0; i < a; i++ {
		go fmt.Println(i)
	}
}

