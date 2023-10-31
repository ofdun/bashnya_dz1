package main

import "fmt"

func main() {
	var n int

	fmt.Println("Введите длину матрицы: ")
	_, err := fmt.Scanf("%d\n", &n)
	if err != nil {
		println(err)
		return
	}

	matrix := make([][]int, n)

	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	fmt.Println("Введите элементы строки матрицы через пробел: ")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			_, err = fmt.Scanf("%d", &matrix[i][j])
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		fmt.Scanf("\n")
	}
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if matrix[i][j] != matrix[j][i] {
				fmt.Println("no")
				return
			}
		}
	}

	fmt.Println("yes")
}
