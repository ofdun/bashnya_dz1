package main

import "fmt"

func main() {
	var n int

	_, err := fmt.Scan(&n)
	if err != nil {
		panic(err)
	}

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		_, err = fmt.Scan(&arr[i])
		if err != nil {
			panic(err)
		}
	}

	tmp := arr[n-1]
	for i := n - 1; i > 0; i-- {
		arr[i] = arr[i-1]
	}
	arr[0] = tmp

	for _, c := range arr {
		fmt.Printf("%d ", c)
	}
}
