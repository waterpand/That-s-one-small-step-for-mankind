package main //Задача. В массиве, состоящем из положительных и отрицательных чисел, определить, сколько элементов превосходят по модулю максимальный элемент.
import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var (
	m               []int
	x, y, z, k, max int
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
	x = 50
	y = 35
	z = 125
	m = slCr(x, y, z)
	fmt.Println("Исходный массив: ")
	fmt.Println(m)
	for i := range m {
		if m[i] > max {
			max = m[i]
		}
	}
	for i := range m {
		if math.Abs(float64(m[i])) > float64(max) {
			k++
		}
	}
	fmt.Println("Максимальный элемент массива: ", max)
	fmt.Println("Количество элементов превосходящих по модулю ", max, " равно: ", k)
}
