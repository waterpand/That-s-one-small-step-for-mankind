package main		// Найти сумму нечетных цифр числа          (a % b) - остаток от деления

import "fmt"
var n, d, sum int
func main(){
	fmt.Println("Нахождение суммы нечетных цифр числа.")
	fmt.Print("Введите целое число:")
	fmt.Scanln(&n)
	for ; n > 0 ; n = n/10 {
		d = (n % 10)
		if ((d % 2) != 0) {
			sum = sum + d
		}
		// fmt.Println("Число =", n)
		// fmt.Println("Остаток =", d)
		// fmt.Println("Сумма =", sum)
	}
	fmt.Println("Сумма =", sum)
}
