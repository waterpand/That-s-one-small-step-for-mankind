package main                            // 1)добавить ввод х и 2)реализовать измененние точности при сравнении

import(
  "fmt"
  "math"
)

func Sqrt(x float64) float64 {
  v := math.Sqrt(x)
  return v
}

func main() {
  var z, z1, x, n float64 = 11, 0, 18, 1
//  n := 1
  i := 0
  fmt.Print("Введите х: ")
  fmt.Scanln(&x)
  fmt.Print("Введите порядок точности для сравнения с функцией math.Sqrt (3..12): ")
  fmt.Scanln(&n)
  s := math.Pow(10, n)
  fmt.Println(float64(s))
  z0 := Sqrt(x)
  for ; (z/z0)-1 > (1/s); z = z1 {            // 1/(float64(10^n)))
    z1 = z - (z * z - x) / (2 * z)
    i++
    fmt.Println(i,") ",z1, "       отношение z/z0 = ", z/z0)
  }
  fmt.Println("    ",z0, " - Квадратный корень math.Sqrt(x)")
}
