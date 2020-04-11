package main //Задача. Определить индексы элементов массива, значение которых лежит в указанном пределе.

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	m           [50]int
	s           []int
	max, min, n int
)

func arrCr() [50]int {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 50; i++ {
		m[i] = rand.Intn(100)
	}
	return m
}

func main() {
	fmt.Println("Определить индексы элементов массива, значение которых лежит в указанном пределе.")
	m = arrCr()
	fmt.Println("Укажите минимальное значение интервала 0...100: ")
	fmt.Scan(&min)
	fmt.Println("Укажите максимальное значение интервала 0...100: ")
	fmt.Scan(&max)
	for i := 0; i < len(m); i++ {
		if m[i] > min && m[i] < max {
			s = append(s, i)
			n++
		}
	}
	fmt.Println("В указанный диапазо попадают ", n, "элементов массива")
	fmt.Println("Их индексы: ", s)
	fmt.Println("Сам массив: ", m)
}

// func arrCr() [25]int {
// 	rand.Seed(time.Now().UnixNano())
// 	for i := 0; i < 25; i++ {
// 		m[i] = rand.Intn(99) - 50
// 	}
// 	return m
// }
