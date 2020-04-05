package main		// Найти сумму четных и нечетных цифр числа          (a % b) - остаток от деления

import "fmt"
var n, d, sum1, sum2 int
func main(){
	fmt.Println("Нахождение сумм четных и нечетных цифр числа.")
	fmt.Print("Введите целое число:")
	fmt.Scanln(&n)
	for ; n > 0 ; n = n/10 {
		d = (n % 10)
		if ((d % 2) != 0) {
			sum1 = sum1 + d
		} else {
			sum2 = sum2 + d
		}
	}
	fmt.Println("Сумма четных цифр числа:", sum2)
	fmt.Println("Сумма нечетных цифр числа:", sum1)
}
