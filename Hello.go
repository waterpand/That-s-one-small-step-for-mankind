package main //Задача. Поменять максимальный и минимальный элементы массива

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	m                 [15]int
	max, min, k, n, p int
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 15; i++ {
		m[i] = rand.Intn(99)
	}
	min = 100
	for i := 0; i < 15; i++ {
		if max < m[i] {
			max = m[i]
			k = i
		}
		if min > m[i] {
			min = m[i]
			n = i
		}
	}
	fmt.Println(min, n)
	fmt.Println(max, k)
	fmt.Println(m)
	p = m[k]
	m[k] = m[n]
	m[n] = p
	fmt.Println(m)
}
