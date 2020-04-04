package main

import(
  "fmt"
)

func main() {
  var (
    z, z1, x float64 = 10, 0, 18
    )
  for i := 0; i < 10; i++ {
    z1 = z - (z * z - x) / (2 * z)
    z = z1
    fmt.Println(z1)
  }
}
