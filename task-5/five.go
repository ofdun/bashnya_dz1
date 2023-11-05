package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string

	_, err := fmt.Scanln(&str)
	if err != nil {
		panic(err)
	}

	result := strings.ReplaceAll(str, "1", "one")

	fmt.Println(result)
}
