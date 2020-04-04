package main

import(
  "fmt"
  "math"
)

func Sqrt(x float64) float64 {
  v := math.Sqrt(x)
  return v
}

func main() {
  var z, z1, x float64 = 10, 0, 18
  z0 := Sqrt(x)
  for i := 0; i < 10; i++ {
    z1 = z - (z * z - x) / (2 * z)
    z = z1
    fmt.Println(z1)
  }
  fmt.Println("Квадратный корень из ", x, " = ", z0)
}
