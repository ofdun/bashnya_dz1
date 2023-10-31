package main

import "fmt"

func main() {
	var n int

	fmt.Println("Введите длину массива: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(err)
		return
	}

	arr := make([]int, n)

	fmt.Printf("Введите %d чисел через пробел: ", n)
	for i := 0; i < n-1; i++ {
		_, err := fmt.Scan(&arr[i+1])
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Scan(&arr[0])

	for _, c := range arr {
		fmt.Printf("%d ", c)
	}
}
