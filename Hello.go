package main //Задача. В одномерном массиве удалить все четные элементы и оставить только нечетные.

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	m    [25]int
	s    []int
	x, y int
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
		if m[i]%2 != 0 {
			s = append(s, m[i])
		}
	}
	fmt.Println(m)
	fmt.Println(s)
}
