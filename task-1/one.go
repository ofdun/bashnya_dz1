package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64

	_, err := fmt.Scanf("%d\n", &a)
	if err != nil {
		panic(err)
	}
	_, err = fmt.Scanf("%d\n", &b)
	if err != nil {
		panic(err)
	}
	fmt.Println(math.Hypot(a, b))
}
