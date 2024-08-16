package basicSolutionAmountAlimony

import (
	"fmt"
)

func Amount() {
	var sumChidlrens int
	var overduePayments string
	var isExpired bool
	var income float64

	var result int
	fmt.Println("Введите количество детей")
	fmt.Scan(&sumChidlrens)
	fmt.Println("Имеется ли просрочка? (Да/Нет)")
	fmt.Scan(&overduePayments)
	fmt.Println("Сумма вашего дохода")
	fmt.Scan(&income)

	if overduePayments == "Да" {
		isExpired = true
	} else {
		isExpired = false
	}

	result = amountAlimony(sumChidlrens, isExpired, income)

	fmt.Printf("Размер алиментов составляет %d", result)
}

func amountAlimony(sumChidlrens int, isExpired bool, income float64) int {
	if isExpired {
		return int(income / 100 * 70)
	}

	var result float64

	if sumChidlrens == 1 {
		result = income / 100 * 25
	} else if sumChidlrens == 2 {
		result = income / 100 * 33
	} else {
		result = income / 100 * 50
	}

	return int(result)
}
