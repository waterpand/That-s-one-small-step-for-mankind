package main //Задача. Определить из каких цифр состоит число

//(a % b) - остаток от деления

import (
	"fmt"
)

var a, b int

func main() {
	fmt.Print("введите число")
	fmt.Scanln(&a)
	for ; a > (99 / 100); a = a / 10 {
		b = b*10 + (a % 10)

	}

	for ; b > (99 / 100); b = b / 10 {
		fmt.Print((b % 10), ",")
	}

}
