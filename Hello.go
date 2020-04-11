package main //Задача. Среднее арифметическое всех чётных элементов массива, стоящих на нечётных местах
import (
	"fmt"
	"math/rand"
	"time"
)

var (
	m     [25]int
	av, n int
	av_n  float64
)

func arrCr(x, y int) [25]int {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 25; i++ {
		m[i] = rand.Intn(x) - y
	}
	return m
}

func main() {
	x := 99
	y := 50
	m = arrCr(x, y)
	fmt.Println(m)
	for i := 0; i < len(m); i++ {
		if m[i]%2 == 0 && i%2 != 0 {
			av = av + m[i]
			n++
		}
	}

	//fmt.Println(av, n)
	av_n = float64(av) / float64(n)
	// av_f := float64(av)
	// n_f := float64(n)
	// av_n = av_f / n_f
	fmt.Println(av_n)
	//fmt.Println(float64(av) / float64(n))
}
