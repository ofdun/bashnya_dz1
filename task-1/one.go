package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64

	fmt.Println("Введите число a > 0: ")
	fmt.Scanf("%g\n", &a)

	fmt.Println("Введите число b > 0: ")
	fmt.Scanf("%g\n", &b)

	c := calc_hypotenuse(a, b)

	fmt.Printf("Гипотенуза равна: %.5g", c)

}

func calc_hypotenuse(a, b float64) float64 {
	c := math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	return c
}
