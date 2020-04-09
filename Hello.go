package main //Задача. Вывести элементы числового массива, которые больше, чем элементы, стоящие перед ними

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	m [15]int
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 15; i++ {
		m[i] = rand.Intn(99)
	}
	for i := 1; i < 15; i++ {
		if m[i-1] < m[i] {
			fmt.Println(m[i])
		}
	}
	fmt.Println(m)
}
