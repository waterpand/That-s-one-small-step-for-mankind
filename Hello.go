package main //Задача.  Посчитать четные и нечетные цифры числа

import "fmt"

var (
	a, i, b, d int
	c, e       int = 1, 1
	num        [20]int
)

func main() {
	fmt.Print("Введите число ")
	fmt.Scanln(&a)
	for ; a > 99/100; a = a / 10 {
		num[i] = (a % 10)
		i++

	}
	i = i - 1 //компенсация лишнего инкремента i++ в последнем цикле
	// fmt.Println(len(num), cap(num))
	for j := 0; i-j > -1; j++ {
		if j%2 != 0 { //чётное
			d = d + num[i-j]
			e = e * num[i-j]
		} else { //нечетное
			b = b + num[i-j]
			c = c * num[i-j]

		}

	}

	fmt.Println("Сумма нечетных ", b, "произведение ", c)
	fmt.Println("Сумма четных ", d, "произведение ", e)
}
