package main

import (
	"fmt"
	"math"
)

//
type Geometry interface {
	area() float64
	perim() float64
}

func measure(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// Cricle
type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perim() float64 {
	return math.Pi * c.radius * 2
}

// Rect
type Rect struct {
	width, height float64
}

func (r *Rect) area() float64 { // haha***
	return r.height * r.width
}

func (r Rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func main() {
	rec := Rect{width: 10.0, height: 10.0}
	fmt.Println(rec, rec.area(), rec.perim())

	rec2 := &rec
	fmt.Println(rec, rec2, rec2.area(), rec2.perim())

	cir1 := Circle{10.0}
	fmt.Println(cir1, cir1.area(), cir1.perim())

	// measure(rec) //error, (r *Rect) area() float64, "r *Rect"
	measure(rec2) // bingo, "rec2 := &rec"
	measure(cir1) // bingo.
}
