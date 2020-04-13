package main //Задача. Найти самые длинные последовательности чисел, упорядоченные по возрастанию.
// через срез срезов реализовать вывод сначала самых длинных последовательностей

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	m                     [80]int
	s                     []int
	t, w, q               [][]int
	x, y, z, max, k, l, c int
)

func arrCr(x, y int) [80]int {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 80; i++ {
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
		for j := 0; j < len(s)-1; j++ {
			if s[j] > s[j+1] {
				if j >= 1 {
					fmt.Println(s[:j+1])
					t = append(t, s[:j+1])
					i = i + j  // первый цикл перепрыгнет отработанный срез
					j = len(s) // выход из второго цикла
				} else {
					i = i + j // происходит тоже самое, что и в случае с true, но срез не выводится на печать, если он состоит из одного элемента
					j = len(s)
				}
			}
		}
	}
	fmt.Println("срез срезов: ", t)
	fmt.Println(len(t))
	c = len(t) - 1
	for l := 0; l < c; l++ {
		k = len(t)
		max = 0
		for i := 0; i < len(t)-1; i++ { // попробовать использовать range  https://go-tour-ru-ru.appspot.com/moretypes/16
			if len(t[i]) > max && i != k {
				max = len(t[i])
				k = i
			}
		}
		fmt.Println(t[k])
		w = append(w, t[k])
		t = append(t[:k], t[k+1:]...)
		fmt.Println("Срез w: ", w)
		fmt.Println("Срез t: ", t)
	}
	w = append(w, t...)
	fmt.Println("Полный Срез w: ", w)
	fmt.Println("Пустой Срез t: ", t)
}
