package main //Задача. Какая сумма элементов массива больше – с первого до элемента с номером К или от элемента с номером К+1 до последнего.
import (
	"fmt"
	"math/rand"
	"time"
)

var (
	m             [25]int
	sum1, sum2, k int
)

func arrCr(x, y int) [25]int {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 25; i++ {
		m[i] = rand.Intn(x) - y
	}
	return m
}

func main() {
	fmt.Print("k= ")
	fmt.Scanln(&k)
	x := 99
	y := 0
	m = arrCr(x, y)
	fmt.Println(m)
	for i := 0; i < len(m); i++ {
		if i <= k {
			sum1 = sum1 + m[i]
		} else {
			sum2 = sum2 + m[i]
		}
	}
	if sum1 > sum2 {
		fmt.Println("Сумма элементов массива с первого и до k-того элементов больше, чем сумма элементов массива с k+1 и до последнеого элемента", sum1, " > ", sum2)
	} else {
		fmt.Println("Сумма элементов массива с первого и до k-того элементов меньше, чем сумма элементов массива с k+1 и до последнеого элемента", sum1, " < ", sum2)
	}

}
