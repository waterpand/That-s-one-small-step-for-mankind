package main // Вывести на экран кубы чисел от A до B, которые вводит пользователь.          (a % b) - остаток от деления

import "fmt"

var (
	a, b int
)

func main() {
	fmt.Print("Введите число А:")
	fmt.Scanln(&a)
	fmt.Print("Введите число B:")
	fmt.Scanln(&b)
	if a > b {
		fmt.Println("Число А должно быть меньше числа В")
	} else {
		for ; a <= b; a++ {
			c := (a * a * a)
			fmt.Println(c)
		}
	}
}
