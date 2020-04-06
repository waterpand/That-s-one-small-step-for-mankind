package main //Задача 2. Найти сумму и количество элементов последовательности, которые по модулю больше 0.001.
//S = 1/2 - 2/4 + 3/8 - 4/16 + ... - ...
//(a % b) - остаток от деления

import (
	"fmt"
	"math"
)

var i int = 1
var a, b, sum, z float64 = 1, 2, 0, 1

func main() {
	for math.Abs(float64(a/b)) > 0.001 {
		sum = sum + (a/b)*z
		z = -z
		fmt.Print("step", i, "   ")
		fmt.Print("a=", a)
		fmt.Print("   b=", b)
		fmt.Println("    sum= ", sum)
		i++
		a = (math.Abs(a) + 1)
		b = b * 2

	}
	fmt.Println("Сумма первых", i, "элементов последовательности равна", sum)
}
