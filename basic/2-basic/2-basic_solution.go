package main
import "fmt"

func main() {
	var a float64

	fmt.Println("Введите значение стороны квадрата: ")
	fmt.Scan(&a)

	fmt.Println("Площадь квадрата: " + fmt.Sprint(a*a))
}

