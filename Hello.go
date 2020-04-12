package main //Задача. Найти два максимальных элемента массива.

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	m                   [25]int
	x, y, max1, max2, k int
)

func arrCr(x, y int) [25]int {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 25; i++ {
		m[i] = rand.Intn(x) - y
	}
	return m
}

func main() {
	x = 99
	y = 50
	m = arrCr(x, y)
	for i := 0; i < len(m); i++ {
		if m[i] > max1 {
			max1 = m[i]
			k = i
		}
	}
	for i := 0; i < len(m); i++ {
		if m[i] > max2 && i != k {
			max2 = m[i]
		}
	}
	fmt.Println(m)
	fmt.Println(max1, max2)
}
