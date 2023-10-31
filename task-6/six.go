package main

import "fmt"

func main() {
	var num, length int
	hashset := make(map[int]struct{})

	fmt.Println("Введите список целых чисел (разделенных пробелами): ")

	for {
		n, err := fmt.Scanf("%d", &num)
		if err != nil {
			break
		}
		if n > 0 {
			if _, exists := hashset[num]; !exists {
				hashset[num] = struct{}{}
				length += 1
			}
		}
	}

	fmt.Println(length)
}
