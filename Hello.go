package main //Задача. Максимальный из отрицательных элементов поменять местами с последним элементом массива.

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	m               [80]int
	max, j, n, x, y int
)

func arrCr(x, y int) [80]int {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 80; i++ {
		m[i] = rand.Intn(x) - y
	}
	return m
}

func main() {
	x = 25
	y = 15
	m = arrCr(x, y)
	fmt.Println("Исходный массив: ")
	fmt.Println(m)
	max = -100
	for i, v := range m {
		if v < 0 && v > max {
			j = i
			max = v
		}
	}
	fmt.Println(max, j)
	if j != len(m)-1 {
		m[j] = m[len(m)-1]
		m[len(m)-1] = max
	}
	fmt.Println("Переработанный массив: ")
	fmt.Println(m)
}
