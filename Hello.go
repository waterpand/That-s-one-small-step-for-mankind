package main //Задача.  Найти сумму цифр ab + cd числа abcd

//Посчитать четные и нечетные цифры числа

import "fmt"

var (
	a, ab, cd int
	num       [4]int
)

func main() {
	fmt.Print("Введите любое четырехзначное число ")
	fmt.Scanln(&a)
	for (a < 1000) || (a > 9999) {
		fmt.Print("четырехзначное число ")
		fmt.Scanln(&a)
	}
	for i := 3; a > 99/100; a = a / 10 {
		num[i] = (a % 10)
		i = i - 1
	}
	ab = num[0]*10 + num[1]
	cd = num[2]*10 + num[3]
	fmt.Println("Сумма цифр ab + cd равна ", ab, "+", cd, "=", ab+cd)
	fmt.Println("произведение ab * cd равно ", ab, "*", cd, "=", ab*cd)
}
