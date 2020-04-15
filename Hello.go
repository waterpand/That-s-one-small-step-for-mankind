package main //Задача. Определить сумму покупки.
// Решение через карты и структуры.

import (
	"fmt"
)

type Goods struct {
	title string
	price float64
}

var (
	priceList = map[int]Goods{
		1: Goods{"Огурцы", 77.50}, 2: Goods{"Помидоры", 180}, 3: Goods{"Яблоки", 95.70}, 4: Goods{"Лимоны", 233}, 5: Goods{"Апельсины", 99.99},
		6: Goods{"Бананы", 67}, 7: Goods{"Мандарины", 117}, 8: Goods{"Груши", 200}, 9: Goods{"Курица", 235.70}, 10: Goods{"Сыр", 473.85},
		11: Goods{"Укроп", 23}, 12: Goods{"Петрушка", 18}, 13: Goods{"Масло", 178.40}, 14: Goods{"Молоко", 67}, 15: Goods{"Сахар", 54},
		16: Goods{"Чай", 80}, 17: Goods{"Виноград", 145},
	}
	weight, sum float64
	j           int = 20
)

func main() {
	for i := 1; i < 18; i++ {
		fmt.Println(i, "--", priceList[i])
	}
	for j != 0 {
		fmt.Print("Введите номер товара или 0 для завершения покупок: ")
		fmt.Scanln(&j)
		if j != 0 {
			p := priceList[j]
			fmt.Print("Вы выбрали ", p.title, " Введите необходимый вес или количество: ")
			fmt.Scanln(&weight)
			sum = sum + weight*p.price
		}
	}
	fmt.Println("Всего покупок на ", sum, "руб.")
	// for i, v := range goods {
	// 	fmt.Println(i+1, "- ", v, "- ", price[i])
	// }
	// for j != 0 {
	// 	fmt.Print("Введите номер товара или 0 для завершения покупок: ")
	// 	fmt.Scanln(&j)
	// 	if j != 0 {
	// 		fmt.Print("Вы выбрали ", goods[j-1], " Введите необходимый вес или количество: ")
	// 		fmt.Scanln(&weight)
	// 		sum = sum + weight*price[j-1]
	// 	}
	// }
	// fmt.Println("Всего покупок на ", sum, "руб.")
}
