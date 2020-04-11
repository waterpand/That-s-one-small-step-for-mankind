package main //Задача. Вывести в порядке возрастания цифры, из которых состоит число и удалить повторяющиеся элементы массива.

import (
	"fmt"
)

var (
	c, min, j, h int
	s, p         []int
)

func main() {
	fmt.Println("Задача: Вывести в порядке возрастания цифры, из которых состоит число.")
	fmt.Print("Введите число: ")
	fmt.Scanln(&c)
	for ; c > 0; c = c / 10 {
		s = append(s, (c % 10))
	}
	fmt.Println(s)
	for i := len(s) - 1; i >= 0; i-- {
		min = 11
		for k := i; k >= 0; k-- {
			if s[k] < min {
				min = s[k]
				j = k
			}
		}
		s = append(s[:j], s[j+1:]...)
		p = append(p, min)
	}
	fmt.Println(p)

	for i := 0; i < len(p); i++ {
		h = p[i]
		for k := 0; k < len(p); k++ {
			if k != i && h == p[k] {
				p = append(p[:k], p[k+1:]...)
				k = 0
			}
		}
	}
	fmt.Println(p)
}
