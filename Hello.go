package main //Задача. Из натурального числа удалить заданную цифру. Число и цифру вводить с клавиатуры.

import (
	"fmt"
)

var a, b, c, d, ch int

func main() {
	fmt.Print("Введите число ")
	fmt.Scanln(&a)
	fmt.Print("Введите цифру, которую следует удалить из числа ")
	fmt.Scanln(&b)
	fmt.Println(a, "исходное число")
	for ; a > (99 / 100); a = a / 10 {
		c = a % 10
		if c != b {
			d = d*10 + c
		}
	}
	for ; d > 99/100; d = d / 10 {
		ch = ch*10 + (d % 10)
	}
	fmt.Println(ch, "число с удаленной цифрой", b)
}
