package main

import (
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("recovered:", recover())
	}()

	defer panic(3)
	defer panic(2)
	defer panic(1)
	panic(0)
}
