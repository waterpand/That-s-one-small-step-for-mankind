package main //Задача. Вывести в порядке возрастания цифры, из которых состоит число

import (
	"fmt"
)

var (
	c, min, j int
	s, p      []int
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
		//fmt.Println(min)
		s = append(s[:j], s[j+1:]...)
		//fmt.Println(s)
		p = append(p, min)
	}
	fmt.Println(p)
}
