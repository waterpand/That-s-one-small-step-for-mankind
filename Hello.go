package main //Задача. Определить сумму покупки.
// Решение через карты и структуры.
//прикрутить добавление или удаление товара.
//убрать, что можно, в функции.
//Сделать, чтобы при выборе товар исчезал из списка.
//Как узнать размер map? --> len(priceList) <--
//Предусмотреть ошибочный ввод номера товара  --> elem, ok = m[key] <--
// Группы товаров..
//При удалении элемента из tempList он так же удаляется из priceList, как этого избежать???

import (
	"fmt"
)

type Goods struct {
	title string
	price float64
}

var (
	priceList = map[int]Goods{
		1: {"Огурцы", 77.50}, 2: {"Помидоры", 180}, 3: {"Яблоки", 95.70}, 4: {"Лимоны", 233}, 5: {"Апельсины", 99.99},
		6: {"Бананы", 67}, 7: {"Мандарины", 117}, 8: {"Груши", 200}, 9: {"Курица", 235.70}, 10: {"Сыр", 473.85},
		11: {"Укроп", 23}, 12: {"Петрушка", 18}, 13: {"Масло", 178.40}, 14: {"Молоко", 67}, 15: {"Сахар", 54},
		16: {"Чай", 80}, 17: {"Виноград", 145},
	}
	tempList, t map[int]Goods
	weight, sum float64
	j           int = 20
)

func listP(tempList map[int]Goods, d int) {
	fmt.Println("В наличии ", len(tempList), "наименований товара: ")
	for i := 1; i <= len(tempList)+d; i++ { //После каждого удаления элемента в функцию возвращается карта с меньшим len и список выводится не весь
		o, ok := tempList[i]
		if ok == true {
			fmt.Println(i, "--", o)
		}
	}
}

func createMap(h map[int]Goods) map[int]Goods {
	t = make(map[int]Goods)
	for k, v := range h {
		t[k] = v
	}
	return t
}

func main() {
	d := 0
	tempList = createMap(priceList)
	for j != 0 {
		listP(tempList, d)
		fmt.Print("Введите номер товара или 0 для завершения покупок: ")
		fmt.Scanln(&j)
		p, ok := tempList[j]
		if j != 0 && j <= len(tempList)+d && ok == true {
			//p := tempList[j]
			fmt.Print("Вы выбрали ", p.title, " Введите необходимый вес или количество: ")
			fmt.Scanln(&weight)
			sum = sum + weight*p.price
			delete(tempList, j)
			d = d + 1 // Не знаю как сделать иначе// счетчик удалений, компенсирует уменьшение длины массива, из-за чего список выводится не весь.
		} else if j > len(tempList)+d && j != 99 {
			fmt.Println("Введен неверный номер товара! Попробуйте снова")
		} else if j == 99 { // этот цикл для проверки первоначальной карты
			for q := 1; q <= len(priceList); q++ {
				om, okk := priceList[q]
				if okk == true {
					fmt.Println(q, "-+-", om)
				} else {
					fmt.Println("Такого ключа не существует")
				}
			}

		} else if j != 0 && j < len(tempList)+d {
			f := priceList[j]
			fmt.Println("Вы уже взяли", f.title, "выберите другой товар или 0 для завершения")
		}
	}
	fmt.Println("Всего покупок на ", sum, "руб.")
}
