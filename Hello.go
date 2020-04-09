package main //Задача. Разделить элементы массива на максимальный

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	m   [15]float64
	max float64
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 15; i++ {
		m[i] = float64(rand.Intn(99))
	}
	for i := 0; i < 15; i++ {
		if max < m[i] {
			max = m[i]
		}
	}
	fmt.Println(max)
	fmt.Println(m)
	for i := 0; i < 15; i++ {
		m[i] = m[i] / max
	}
	fmt.Println(m)
}
