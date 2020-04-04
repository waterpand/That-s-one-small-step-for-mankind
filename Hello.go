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
  var z, z1, x float64 = 11, 0, 18
  i := 0
  z0 := Sqrt(x)
  for ; ((z/z0)-1 > 0.0000000000001); z = z1 {
    z1 = z - (z * z - x) / (2 * z)
    i++
    fmt.Println(i,") ",z1, "       отношение z/z0 = ", z/z0)
  }
  fmt.Println(z0, " - Квадратный корень math.Sqrt(x)")
}
