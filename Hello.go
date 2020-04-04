package main          // не работает, если корень из х больше предустановленного z (сейчас 999999)

import(
  "fmt"
  "math"
)

func Sqrt(x float64) float64 {
  v := math.Sqrt(x)
  return v
}

func main() {
  var z, z1, x, n float64 = 999999, 0, 18, 1
  i := 0
  fmt.Print("Введите х: ")
  fmt.Scanln(&x)
  fmt.Print("Введите порядок точности для сравнения с функцией math.Sqrt (3..12): ")
  fmt.Scanln(&n)
  s := math.Pow(10, n)
  z0 := Sqrt(x)
  for ; (z/z0)-1 > (1/s); z = z1 {
    z1 = z - (z * z - x) / (2 * z)
    i++
  }
  fmt.Println("Выполнено ", i," итерации(й) вычисления.")
  fmt.Println(z1, " -- значение, полученное методом Ньютона")
  fmt.Println(z0, " -- Квадратный корень, вычесленный функцией math.Sqrt(x)")
}
