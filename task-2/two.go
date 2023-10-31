package main

import "fmt"

func main() {
	var triangle bool

	var a, b, c int16

	fmt.Println("Введите первую сторону: ")
	_, err := fmt.Scanf("%d\n", &a)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Введите вторую сторону: ")
	_, err = fmt.Scanf("%d\n", &b)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Введите третью сторону: ")
	_, err = fmt.Scanf("%d\n", &c)
	if err != nil {
		fmt.Println(err)
		return
	}

	triangle = isTriangle(a, b, c)

	if triangle {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func isTriangle(a, b, c int16) bool {

	if a < 0 || b < 0 || c < 0 {
		return false
	}
	if a > b+c || c > a+b || b > a+c {
		return false
	}

	return true
}
