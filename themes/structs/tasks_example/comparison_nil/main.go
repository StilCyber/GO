package main

import "fmt"

type Data struct{}

func (d *Data) Print() {
	if d == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}

func main() {
	var data *Data
	data.Print()
}