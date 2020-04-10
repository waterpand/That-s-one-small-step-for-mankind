package main //Задача. Заменить элементы массива на противоположные

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	m          [20]int
	sum1, sum2 int
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ { // Заполнение массива
		m[i] = rand.Intn(99) - 50
		sum1 = sum1 + m[i]
	}
	fmt.Println(m, "  sum1 = ", sum1)
	for i := 0; i < 20; i++ {
		m[i] = -m[i]
		sum2 = sum2 + m[i]
	}
	fmt.Println(m, "  sum2 = ", sum2)
}
