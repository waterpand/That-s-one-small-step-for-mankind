package main //Задача. В однородном массиве, состоящем из N вещественных элементов, найти максимальный по модулю элемент массива..

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	m         [25]int
	max, n, k int
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

	for i := 0; i < 25; i++ {
		n = m[i]
		if n < 0 {
			n = -n
		}
		if n > max {
			max = n
			k = i
		}
	}
	fmt.Println(m)
	fmt.Println(m[k])
}
