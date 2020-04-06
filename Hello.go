package main //Вычисление факториала числа             после 20 значение факториала становится отрицательным, почему?
//(a % b) - остаток от деления

import "fmt"

var a int
var f int = 1

func main() {
	fmt.Print("Введите число :")
	fmt.Scanln(&a)

	for i := 1; i <= a; i++ {
		f = f * i
		fmt.Println(f)
	}
	fmt.Println("Факториал числа", a, "равен", f)
}
