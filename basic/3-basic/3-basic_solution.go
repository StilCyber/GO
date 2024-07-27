package main 
import ("fmt"; "math/rand")

func main() {
	var randomNumber float64 = rand.Float64()
	if randomNumber > 0.5 {
		fmt.Println(randomNumber)
		fmt.Println("Лучше взять зонт")
	} else {
		fmt.Println(randomNumber)
		fmt.Println("Можно с зонтом не таскаться")
	}
}