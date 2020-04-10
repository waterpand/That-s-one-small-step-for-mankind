package main //Задача. Вывести элементы массива, которые больше среднего арифметического

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	m   [25]int
	avr int
)

func ArrCr() [25]int {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 25; i++ {
		m[i] = rand.Intn(155)
	}
	return m
}

func main() {
	m = ArrCr()
	//fmt.Println(m)
	for i := 0; i < len(m); i++ {
		avr = avr + m[i]
	}
	avr = avr / 25
	fmt.Println("Среднеарифметическое массива: ", avr)
	avrrr := []int{}
	for i := 0; i < len(m); i++ {
		if avr < m[i] {
			//fmt.Print(m[i], " ")
			avrrr = append(avrrr, m[i])
		} else {
			avrrr = append(avrrr, 0)
		}
	}
	//fmt.Println()
	fmt.Println(m)
	fmt.Println(avrrr)
}
