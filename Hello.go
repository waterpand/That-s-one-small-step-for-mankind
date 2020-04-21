package main //Задача. Найти максимальный элемент численного массива
import (
	"fmt"
	"math/rand"
	"time"
)

var (
	m            []int
	x, y, z, max int
)

func sliceCreate(x, y, z int) []int {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < z; i++ {
		m = append(m, rand.Intn(x)-y)
	}
	return m
}

func main() {
	fmt.Print("Slice size: ")
	fmt.Scanln(&z)
	x = 99
	m = sliceCreate(x, y, z)
	fmt.Println(m)
	for _, i := range m {
		if i > max {
			max = i
		}
	}
	fmt.Println(max)
}
