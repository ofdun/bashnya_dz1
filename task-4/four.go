package main

import "fmt"

func main() {
	var n int

	_, err := fmt.Scanf("%d\n", &n)
	if err != nil {
		panic(err)
	}

	matrix := make([][]int, n)

	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			_, err = fmt.Scanf("%d", &matrix[i][j])
			if err != nil {
				panic(err)
			}
		}
		fmt.Scanf("\n")
	}
	if isSymmetricMatrix(matrix) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}

func isSymmetricMatrix(matrix [][]int) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] != matrix[j][i] {
				return false
			}
		}
	}
	return true
}
