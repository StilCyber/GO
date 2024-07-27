package main 
import "fmt"

func main() {
	var str string 
	fmt.Println("Если изменила жена, введите - Yes, иначе введите No")
	fmt.Scan(&str)

	switch str {
	case "Yes":
		fmt.Println("Муж сам виноват, не уделял внимание, нужно забрать у него все") 
	case "No": 
		fmt.Println("Муж виноват - изменщика нужно наказать")
	}
}