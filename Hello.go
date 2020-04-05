package main

import(
  "fmt"
  )

func main() {
    var i int = 0
    fmt.Print("Введи число 1..5: ")
    fmt.Scanln(&i)
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
    case 4:
        fmt.Println("foure")
    case 5:
        fmt.Println("five")
	default:
		fmt.Println("wrong number!")
	}
}
