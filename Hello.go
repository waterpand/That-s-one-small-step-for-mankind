package main //Задача. Определение различных (не повторяющихся) цифр, входящих в число

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	m    []int
	c, n int
)

func main() {
	fmt.Println("Введите произвольный множитель числа: ")
	fmt.Scan(&n)
	rand.Seed(time.Now().UnixNano())
	c = (rand.Intn(9999) * 9777) * n
	fmt.Println(c, " -- полученное число")
	m = append(m, (c % 10))
	for ; c > 0; c = c / 10 {
		f := false
		for j := 0; j < len(m); j++ {
			if m[j] == (c % 10) {
				f = true
			}
		}
		if f == false {
			m = append(m, (c % 10))
		}
	}
	for i := len(m) - 1; i >= 0; i-- {
		fmt.Print(m[i], ", ")
	}
}
