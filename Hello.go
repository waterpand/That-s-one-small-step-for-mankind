package main //Задача. Вывести на экран значения массива по спирали. Будем понимать под этим следующее. Если массив состоит из 100 элементов, то выводить по 10 элементов в каждой строке, при этом каждая вторая строка должна выводиться "задом наперед": сначала последние элементы строки, а затем первые.
import (
	"fmt"
	"math/rand"
	"time"
)

var (
	m       []int
	z, y, x int
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
	x = 1000
	y = 0
	fmt.Println("Введите размер массива: ")
	fmt.Scan(&z)
	m = slCr(x, y, z)
	for i := 0; i < z; i++ {
		if (i/10)%2 != 0 && i%10 == 9 {
			fmt.Println("   ", m[i])
		} else if (i/10)%2 != 0 {
			fmt.Print("   ", m[i])
		} else if (i/10)%2 == 0 && i%10 == 9 {
			fmt.Println("   ", m[i])
		} else if (i/10)%2 == 0 {
			fmt.Print("   ", m[i])
		}
	}
	fmt.Println()
	for i := 0; i < z; i = i + 10 {
		switch (i/10)%2 == 0 {
		case true:
			fmt.Print("---> ")
			for k := i; k < i+10; k++ {
				if k < len(m) {
					fmt.Print(m[k], " ")
				}
			}
			fmt.Println()
		case false:
			fmt.Print("<--- ")
			for k := i + 9; k >= i; k-- {
				if k < len(m) {
					fmt.Print(m[k], " ")
				}
			}
			fmt.Println()
		}
	}
}
