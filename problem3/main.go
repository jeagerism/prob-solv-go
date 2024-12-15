package main

import "fmt"

func main() {
	rec := new(Rectangle)
	rec.height = 5
	rec.width = 5

	fmt.Printf("Area : %f And Perimeter : %f", rec.Area(), rec.Perimeter())

}

type Rectangle struct {
	width  float32
	height float32
}

func (r *Rectangle) Area() float32 {
	return r.height * r.width
}

func (r *Rectangle) Perimeter() float32 {
	return r.height*2 + r.width*2
}
