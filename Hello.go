package main

import (
  "fmt"
)

func main() {
  var a int
  fmt.Println("Hello, world", "Введи число: ")
  fmt.Scan(&a)
  if a < 5 {
    fmt.Println("Сопляк!")
  } else if a == 5 {
      fmt.Println("В детский сад, щегол!")
  } else if (a > 5) && (a < 18) {
      fmt.Println("Почему не в школе, шкед?")
      fmt.Println("Дуй в ", (a - 6), " класс!" )
  } else if (a >= 18) && (a <= 23) {
      fmt.Println("К сессии готов?")
  } else if (a > 23) && (a <= 27) {
      fmt.Println("Становись, ровняйсь, смирно!")
  } else if (a > 27) {
      fmt.Println("да не парься уже, до пенсии осталось всего ", (65 - a), " лет. Посторайся дотянуть...")
  }
}
