package main

import "fmt"

func main() {
	var num int
	hashset := make(map[int]struct{})

	for {
		_, err := fmt.Scanf("%d", &num)
		if err != nil {
			break
		}
		if _, exists := hashset[num]; !exists {
			hashset[num] = struct{}{}
		}
	}

	fmt.Println(len(hashset))
}
