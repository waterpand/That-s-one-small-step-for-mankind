package main //Задача.   ребуется определить, сколько можно преобрести ручек (по цене 10 руб.), карандашей (5 руб.) и ластиков (2 руб.) на 100 рублей. При этом всего предметов должно быть 30.

import "fmt"

var (
	x, y, z, i int
	a, b, c    int = 10, 5, 2
)

func main() {
	fmt.Println("Введите A, B, C")
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)
	for x := 0; x < 30; x++ {
		for y := 0; y < 30; y++ {
			for z := 0; z < 30; z++ {
				if (a*x+b*y+c*z == 100) && (x+y+z == 30) {
					i++
					fmt.Println("Решение №", i)
					fmt.Println("X:", x)
					fmt.Println("Y:", y)
					fmt.Println("Z:", z)
					fmt.Println()
				}
			}
		}
	}
	if i == 0 {
		fmt.Println("Уравнение не имеет решений")
	}
}
