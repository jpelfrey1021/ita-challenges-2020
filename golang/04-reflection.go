package main

import (
	"fmt"
	"reflect"
)

//setting up types
type shape interface {
	area() float64
}

type square struct {
	side float64
}

type rectangle struct {
	width  float64
	height float64
}

//main func
func main() {
	rec := rectangle{2, 3}
	sqr := square{5.5}
	fmt.Println(myType(rec))
	fmt.Println(myType(sqr))
}

//individual func
func (s square) area() float64 {
	return s.side * s.side
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

//interface func
func myType(s shape) string {
	bol := "rectangle" == reflect.TypeOf(s).Name()
	if bol {
		return "rectangle"
	}
	return "square"
}
