package main //Задача. Найти самые длинные последовательности чисел, упорядоченные по возрастанию.
// через срез срезов реализовать вывод сначала самых длинных последовательностей

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	m               [80]int
	s               []int
	t, w            [][]int
	x, y, z, max, k int
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
	fmt.Println(len(t), len(t[4]))
	k = len(t)
	for i := 0; i < len(t)-1; i++ {
		if len(t[i]) > max && i != k {
			max = len(t[i])
			k = i
			//w = append(t, t[k])
			//w = append(t, t[k+1])
		}
	}
	fmt.Println(t)
	fmt.Println(t[k])
	fmt.Println("Срез w: ", w)
	//w = append(w, t[z][k])
	//z++
}
