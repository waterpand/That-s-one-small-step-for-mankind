package main //Задача. Найти самые длинные последовательности чисел, упорядоченные по возрастанию.

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	m    [50]int
	s, t []int
	x, y int
)

func arrCr(x, y int) [50]int {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 50; i++ {
		m[i] = rand.Intn(x) - y
	}
	return m
}

func main() {
	x = 16
	y = 0
	m = arrCr(x, y)
	fmt.Println(m)
	for i := 0; i < len(m); i++ {
		s := m[i:] // срез от текущей позиции и до конца массива
		//fmt.Println(s)
		for j := 0; j < len(s)-1; j++ {
			if s[j] > s[j+1] {
				if j >= 1 {
					fmt.Println(s[:j+1])
					i = i + j  // первый цикл перепрыгнет отработанный срез
					j = len(s) // выход из второго цикла
				} else {
					i = i + j // происходит тоже самое, что и в случае с true, но срез не выводится на печать, если он состоит из одного элемента.
					j = len(s)
				}
			}
		}
	}
}
