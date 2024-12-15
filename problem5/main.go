package main

import (
	"fmt"
	"math"
)

func main() {
	rect := &Rectangle{width: 5, height: 10}
	circle := &Circle{radias: 7}

	// ใช้งาน Polymorphism ผ่าน Shape interface
	printShapeInfo(rect)
	printShapeInfo(circle)
}

func printShapeInfo(s Shape) {
	fmt.Printf("Shape Info:\n")
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter : %.2f\n", s.Perimeter())
	fmt.Println()
}

type Shape interface {
	Area() float32
	Perimeter() float32
}

// Rectangle

type Rectangle struct {
	width  float32
	height float32
}

func (r *Rectangle) Area() float32 {
	return r.width * r.height
}

func (r *Rectangle) Perimeter() float32 {
	return 2*r.width + 2*r.height
}

// Circle

type Circle struct {
	radias float32
}

func (c *Circle) Area() float32 {
	return math.Pi * c.radias * c.radias
}

func (c *Circle) Perimeter() float32 {
	return 2 * math.Pi * c.radias
}
