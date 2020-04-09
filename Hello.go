package main //Задача.   Вычислить сумму ряда чисел 1/1^2 + 1/2^2 + 1/3^2 + … + 1/n^2, где n определяется пользователем

import "fmt"

var (
	n, sum float64
)

func main() {
	fmt.Print("Введите число: ")
	fmt.Scanln(&n)
	for ; n != 0; n-- {
		sum = sum + (1 / (n * n))
	}
	fmt.Println("Сумма ряда равна: ", sum)
}
