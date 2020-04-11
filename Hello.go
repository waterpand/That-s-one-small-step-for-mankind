package main //Задача. Дан массив A вещественного типа, содержащий 20 положительных и отрицательных элементов. Сформировать массив B из положительных элементов массива A, имеющих четный индекс. Найти сумму квадратов элементов нового массива.

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	a   [20]int
	b   []int
	sum int
)

func arrCr() [20]int {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		a[i] = rand.Intn(99) - 50
	}
	return a
}

func main() {
	a = arrCr()
	fmt.Println(a)
	for i := 0; i < 20; i++ {
		if a[i] > 0 && i%2 == 0 {
			b = append(b, a[i])
			sum = sum + a[i]*a[i]
		}
	}
	fmt.Println(b)
	fmt.Println(sum)
}
