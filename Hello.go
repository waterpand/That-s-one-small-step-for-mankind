package main //Задача. имеется одномерный массив, содержащий числа от 0 до 49 включительно. Требуется исключить из него все элементы, значения которых меньше 15.

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	m, n       []int
	x, y, z, k int
)

func sliceCreate(x, y, z int) []int {
	//m = make([]int, z)						работает и так и так
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < z; i++ {
		m = append(m, rand.Intn(x)-y)
		//m[i] = rand.Intn(x) - y
	}
	return m
}

func main() {
	x = 50
	y = 0
	fmt.Print("Введите размер массива: ")
	fmt.Scanln(&z)
	m = sliceCreate(x, y, z)
	fmt.Println("Исходный массив: ", m)
	fmt.Println("Исключить из исходного массива все элементы меньше 15")
	for i := 0; i < z; i++ {
		if m[i] >= 15 {
			n = append(n, m[i])
		}
	}
	fmt.Println("обработанный массив: ", n)
}
