package eventLoopBreak

import "fmt"

func EventLoopBreak() bool {

	wife := true
	var honor string

	for wife {
		fmt.Println("Вам еще не надоело? Может, подать на развод? Да/нет")
		fmt.Scan(&honor)
		if(honor == "Да") {
			fmt.Println("Поздравляем со свободой (^_^)")
			break
		}
		
	}

	return wife
}