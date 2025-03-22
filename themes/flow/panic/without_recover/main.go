package main

import (
	"fmt"
)

func process() {
	fmt.Println("V2: open file")
	defer fmt.Println("V2: close file")

	panic("error")
}

// func process() {
// 	fmt.Println("V2: open file")
// 	defer fmt.Println("V2: close file")

// 	runtime.Goexit()
// }

// func process() {
// 	fmt.Println("V2: open file")
// 	defer fmt.Println("V2: close file")

// 	os.Exit(1)
// }

func main() {
	process()
}