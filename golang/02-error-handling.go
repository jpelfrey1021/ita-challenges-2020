package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(stringToInt("9"))
}

func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}
