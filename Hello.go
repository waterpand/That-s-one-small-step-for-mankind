package main //Задача.   найти все комбинации из трех чисел до определенного предела, которые в сумме дают другое число.

import "fmt"

var a, n int

func main() {
	fmt.Print("Введите число: ")
	fmt.Scan(&a)
	for i := 0; i < a; i++ {
		for j := 0; j < a; j++ {
			for k := 0; k < a; k++ {
				if i+j+k == a {
					fmt.Println(i, j, k)
					n++
				}
			}
		}
	}
	fmt.Print("Произведено итераций: ", n+1)
}
