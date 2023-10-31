package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64

	fmt.Println("Введите число a > 0: ")
	_, err := fmt.Scanf("%g\n", &a)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Введите число b > 0: ")
	_, err = fmt.Scanf("%g\n", &b)
	if err != nil {
		fmt.Println(err)
		return
	}

	c := calc_hypotenuse(a, b)

	fmt.Println(c)

}

func calc_hypotenuse(a, b float64) float64 {
	c := math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	return c
}
