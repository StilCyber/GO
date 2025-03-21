package main

import "fmt"

func main() {
	var temp float32
	fmt.Println("Введите температуру в градусах по цельсию")
	fmt.Scan(&temp)
	fmt.Println("Температура в фаренгейтах: " + fmt.Sprint(temp*9/5 + 32))
}