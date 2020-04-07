package main //Задача. Вывести все квадраты натуральных чисел, не превосходящие данного числа N.
//Пример: N=50 | 1 4 9 16 25 36 49

import (
	"fmt"
)

var n, x int

func main() {
	fmt.Print("Введите число ")
	fmt.Scan(&n)
	for i := 1; (i * i) < n; i++ {
		fmt.Println(i, "   ", i*i)
	}
}
