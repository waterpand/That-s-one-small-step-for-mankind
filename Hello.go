package main //Задача. Расстояние между точками в n-мерном пространстве
//При заданных координатах A1, A2, ..., An одной точки и координатах B1, B2, ..., Bn другой точки n-мерного пространства. Найти расстояние между ними по формуле sqrt(sqr(A1-B1) + ... + sqr(An-Bn)).

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var (
	cA, cB []int
	n      int
	d      float64
)

func CreatePoint(int) []int {
	rand.Seed(time.Now().UnixNano())
	coord := []int{}
	for i := 0; i < n; i++ {
		coord = append(coord, rand.Intn(10))
	}
	return coord
}

func main() {
	// fmt.Print("Определить мерность пространства: ")
	// fmt.Scanln(&n)
	rand.Seed(time.Now().UnixNano())
	n = rand.Intn(11)
	fmt.Println("Мерность пространства: ", n)
	cA = CreatePoint(n)
	fmt.Println("Координаты точки А: ", cA)
	time.Sleep(10 * time.Millisecond)
	cB = CreatePoint(n)
	fmt.Println("Координаты точки В: ", cB)

	for i := 0; i < n; i++ {
		d = d + float64((cA[i]-cB[i])*(cA[i]-cB[i]))
	}

	fmt.Println(d)
	d = math.Sqrt(d)
	fmt.Println("Расстояние между точками А и В: ", d)

}
