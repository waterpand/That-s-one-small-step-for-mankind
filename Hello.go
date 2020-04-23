// Playing with digits https://www.codewars.com/kata/5552101f47fc5178b1000050/train/go

// Some numbers have funny properties. For example:

// 89 --> 8¹ + 9² = 89 * 1

// 695 --> 6² + 9³ + 5⁴= 1390 = 695 * 2

// 46288 --> 4³ + 6⁴+ 2⁵ + 8⁶ + 8⁷ = 2360688 = 46288 * 51

// Given a positive integer n written as abcd... (a, b, c, d... being digits) and a positive integer p

// we want to find a positive integer k, if it exists, such as the sum of the digits of n taken to the successive powers of p is equal to k * n.
// In other words:

// Is there an integer k such as : (a ^ p + b ^ (p+1) + c ^(p+2) + d ^ (p+3) + ...) = n * k

// If it is the case we will return k, if not return -1.

// Note: n and p will always be given as strictly positive integers.

package main

import (
	"fmt"
	"math"
)

var (
	m             []int
	n, p, sum, n1 int = 89, 1, 0, 0 // n = 89, n = 92, n = 695
)

func main() {
	n1 = n
	for ; n1 >= 1; n1 = n1 / 10 {
		m = append(m, n1%10)
		//fmt.Println(m)
	}
	for i := len(m); i > 0; i-- {
		sum = sum + int(math.Pow(float64(m[i-1]), float64(p+len(m)-i)))
		//fmt.Println(sum, m[i-1], p, i, p+i)
	}
	if sum%n == 0 {
		fmt.Println("return", sum/n)
	} else {
		fmt.Println("-1")
	}

}
