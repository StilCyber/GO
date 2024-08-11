package basicSeven

import ("fmt")

func CalcFactorial(num int) (int) {
	
	var result int = 1

	for i := 1; i <= num; i++ {
		result *= i
	}

	fmt.Println(result)
	return result
}