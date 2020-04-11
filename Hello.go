package main //Задача. Найти среднее арифметическое отрицательных элементов массива. Заменить на него минимальный элемент.

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	m        [25]int
	av, c, n int
	min      int = 50
)

func arrCr() [25]int {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 25; i++ {
		m[i] = rand.Intn(99) - 50
	}
	return m
}

func main() {
	m = arrCr()
	fmt.Println(m)
	for i := 0; i < len(m); i++ {
		if m[i] < 0 {
			av = av + m[i]
			c++
		}
		if m[i] < min {
			min = m[i]
			n = i
		}
	}
	m[n] = av / c
	fmt.Println("Минимальный элемент массива", min, "с индексом", n, "меняется на среднеарифметическое отрицательных элементов", av/c)
	fmt.Println(m)
}
