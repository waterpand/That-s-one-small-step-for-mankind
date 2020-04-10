package main //Задача. В массиве найти минимальное значение среди элементов с нечетными индексами.

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	m         [20]int
	sum, sum1 int = 50, 0
)

func arrCr() [20]int { // Заполнение массива
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		m[i] = rand.Intn(99) - 50
	}
	return m
}

func main() {
	for k := 0; k < 100; k++ {
		m = arrCr()
		for i := 0; i < len(m); i++ {
			if m[i] < sum && i%2 != 0 {
				sum = m[i]
			}
		}
		fmt.Println(m)
		fmt.Println(sum)
		sum1 = sum1 + sum
	}
	fmt.Println(sum1 / 100)
}
