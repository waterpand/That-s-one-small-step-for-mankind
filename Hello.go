package main //Задача. Вставка элемента в срез
import (
	"fmt"
	"math/rand"
	"time"
)

var (
	m, n             []int
	x, y, z, k, l, p int
)

func sliceCreate(x, y, z int) []int {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < z; i++ {
		m = append(m, rand.Intn(x)-y)
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
	fmt.Print("Введите элемент, который необходимо вставить в срез: ")
	fmt.Scanln(&l)
	fmt.Print("Введите индекс нового элемента: ")
	fmt.Scanln(&k)
	n = append(n[:], m[:k]...)
	n = append(n, l)
	n = append(n[:], m[k:]...)
	fmt.Println(n)
}
