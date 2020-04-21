package main //Задача. Определить количество элементов массива, значение которых больше соседних элементов
import (
	"fmt"
	"math/rand"
	"time"
)

var (
	m          []int
	z, y, x, q int
)

func slCr(x, y, z int) []int {
	m = make([]int, z)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < z; i++ {
		m[i] = rand.Intn(x) - y
	}
	return m
}

func main() {
	x = 100
	y = 0
	fmt.Println("Введите размер массива: ")
	fmt.Scan(&z)
	m = slCr(x, y, z)
	for i := 1; i < z-1; i++ {
		if m[i] > m[i-1] && m[i] > m[i+1] {
			fmt.Print("   ", i, ") - ", m[i])
			q = q + 1
		}
	}
	fmt.Println()
	fmt.Println(q)
}
