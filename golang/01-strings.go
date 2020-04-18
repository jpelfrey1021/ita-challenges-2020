package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Fact is Not a Reason"
	fmt.Println(acronymMaker(s))
}

func acronymMaker(s string) string {
	x := strings.Split(s, " ")
	var w string
	y := []string{}

	for _, v := range x {
		y = append(y, ((strings.Split(v, ""))[0]))
		w = strings.ToUpper(strings.Join(y, ""))

	}

	return w
}
