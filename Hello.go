                                                // a[0:10] эти записи эквивалентны
                                                // a[:10] Значением по умолчанию для нижней границы является нуль, а для верхней - размер среза.
                                                // a[0:]
                                                // a[:]



package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]                                 // [3 5 7]
	fmt.Println(s)

	s = s[:2]                                  // [3 5]
	fmt.Println(s)

	s = s[1:]                                  // [5]
	fmt.Println(s)

    fmt.Println(3*3)
    fmt.Println(1E3)
}
