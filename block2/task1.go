package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	a float64 // ширина
	h float64 //высота
}

type Circle struct {
	r float64 //радиус
}

type Figures interface {
	Area() float64
	Type() string
}

func NewRectangle(a, h float64) *Rectangle {
	return &Rectangle{a: a, h: h}
}

func NewCircle(r float64, c int) *Circle {
	return &Circle{r: r}
}

func (r Rectangle) Area() float64 {
	S := r.a * r.h
	return S
}

func (c Circle) Area() float64 {
	S := math.Pi * c.r
	return S
}

func (r Rectangle) Type() string {
	return "rectangle"
}

func (c Circle) Type() string {
	return "circle"
}

func Print(f Figures) {
	area := f.Area()
	figType := f.Type()
	fmt.Println(area)
	fmt.Println(figType)
}

func main() {
	r := NewRectangle(25.0, 50.0)
	c := NewCircle(2, 2)
	Print(r)
	Print(c)
}
