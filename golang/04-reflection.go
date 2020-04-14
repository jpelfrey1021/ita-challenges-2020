package challenge

import "reflect"

type shape interface {
	area() int
}

type square struct {
	side int
}

type rectangle struct {
	width, height int
}

func (s square) area() int {
	return s.side * s.side
}

func (r rectangle) area() int {
	return r.width * r.height
}

func myType(s shape) string {
	bol := "rectangle" == reflect.TypeOf(s).Name()
	if bol {
		return "rectangle"
	}
	return "square"
}