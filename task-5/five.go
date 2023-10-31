package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string

	fmt.Println("Введите строку: ")

	_, err := fmt.Scanln(&str)
	if err != nil {
		fmt.Println(err)
		return
	}

	result := strings.ReplaceAll(str, "1", "one")

	fmt.Println(result)
}
