package main //Задача. Найти сумму первой и последней цифр любого целого положительного числа.

//Найти сумму цифр ab + cd числа abcd
//Посчитать четные и нечетные цифры числа

import "fmt"

var (
	a, i int
	num  [20]int
)

func main() {
	fmt.Print("Введите число ")
	fmt.Scanln(&a)
	for ; a > 99/100; a = a / 10 {
		num[i] = (a % 10)
		i++
	}
	fmt.Println("Первая цифра числа ", num[i-1])
	fmt.Println("Последняя цифра числа ", num[0])
	fmt.Println("Их сумма ", num[0]+num[i-1])
}
