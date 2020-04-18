package main //Задача. Определить сумму покупки.
//прикрутить добавление или удаление товара.  -->    <--
//убрать, что можно, в функции.
//Сделать, чтобы при выборе товар исчезал из списка. --> delete(tempList, j) <--
//Как узнать размер map? --> len(priceList) <--
//Предусмотреть ошибочный ввод номера товара  --> elem, ok = m[key] <--
// Группы товаров..
// сделать обновление листа для добавленных или удаленных -- не работает, проблема в размере карты

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
	tempList, t     map[int]Goods
	weight, sum, pr float64
	tit             string
	x               int
	j               int = 20
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

func addGoods(tit string, pr float64, z int) int {
	stop := true
	fmt.Println("1", stop, z)
	for i := 1; i <= z; i++ {
		_, ok := priceList[i]
		fmt.Println("2", stop, ok)
		if ok != true && stop == true {
			priceList[i] = Goods{tit, pr}
			stop = false
			fmt.Println("3", stop, i, z)
		}
	}
	fmt.Println("4", stop, z)
	if stop == true {
		key := z + 1
		z = z + 1
		fmt.Println("5", stop, key, z)
		fmt.Println(key)
		priceList[key] = Goods{tit, pr}
		fmt.Println(priceList[key])

	}
	return z
}

func delGoods(x int) {
	delete(priceList, x)
}

func main() {
	z := len(priceList)
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
		} else if j > len(tempList)+d && j != 99 && j != 101 && j != 202 && j != 303 {
			fmt.Println("Введен неверный номер товара! Попробуйте снова")
		} else if j == 99 { // этот цикл для проверки первоначальной карты
			fmt.Println(d, "d")
			for q := 1; q <= z; q++ {
				om, okk := priceList[q]
				if okk == true {
					fmt.Println(q, "-+-", om)
				} else {
					fmt.Println(q, "Данный товар был удален")
				}
			}
		} else if j == 101 { // добавление элемента
			fmt.Println("Добавление товара в список: ")
			fmt.Print("Введите название товара ")
			fmt.Scanln(&tit)
			fmt.Print("Введите цену на ", tit, " ")
			fmt.Scanln(&pr)
			z = addGoods(tit, pr, z)
		} else if j == 202 { // удаление товара
			fmt.Print("Выберите товар, который будет удален из списка: ")
			fmt.Scanln(&x)
			delGoods(x)
			d = d + 1
		} else if j == 303 {
			tempList = createMap(priceList)
		} else if j != 0 && j <= len(tempList)+d {
			f := priceList[j]
			fmt.Println("Вы уже взяли", f.title, "выберите другой товар или 0 для завершения")
		}
	}
	fmt.Println("Всего покупок на ", sum, "руб.")
}
