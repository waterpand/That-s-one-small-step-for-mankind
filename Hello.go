package main //Задача. Найти сумму четных отрицательных элементов массива

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	m   [20]int
	sum int
)

func cycle() [20]int {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		m[i] = rand.Intn(99) - 50
	}
	return m
}

func main() {
	//rand.Seed(time.Now().UnixNano())
	//for i := 0; i < 20; i++ { // Заполнение массива
	//	m[i] = rand.Intn(99) - 50

	//}
	m = cycle()
	fmt.Println(m)
	for i := 0; i < 20; i++ {
		if m[i] < 0 && (m[i]%2) == 0 {
			sum = sum + m[i]
		}
	}
	fmt.Println(m, sum)
}
