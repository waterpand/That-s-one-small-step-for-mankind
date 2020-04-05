package main // Нахождение количества четных и нечетных цифр числа и их сумм          (a % b) - остаток от деления

import "fmt"

var n, d, sum1, sum2, i1, i2 int

func main() {
	fmt.Println("Нахождение количества четных и нечетных цифр числа и их сумм.")
	fmt.Print("Введите целое число:")
	fmt.Scanln(&n)
	for ; n > 0; n = n / 10 {
		d = (n % 10)
		if (d % 2) != 0 {
			sum1 = sum1 + d
			i1++
		} else {
			sum2 = sum2 + d
			i2++
		}
	}
	fmt.Println("Количество четных цифр числа:", i2)
	fmt.Println("Сумма четных цифр числа:", sum2)
	fmt.Println()
	fmt.Println("Количество нечетных цифр числа:", i1)
	fmt.Println("Сумма нечетных цифр числа:", sum1)
}
