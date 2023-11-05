package main

import "fmt"

func main() {
	var a, b, c int

	_, err := fmt.Scanf("%d\n", &a)
	if err != nil {
		panic(err)
	}
	_, err = fmt.Scanf("%d\n", &b)
	if err != nil {
		panic(err)
	}
	_, err = fmt.Scanf("%d\n", &c)
	if err != nil {
		panic(err)
	}

	if isTriangle(a, b, c) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func isTriangle(a, b, c int) bool {
	return a < b+c && c < a+b && b < a+c
}
