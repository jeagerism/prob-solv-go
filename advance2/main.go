package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	// กำหนดรูปทรง
	shapes := []Shape{
		&Rectangle{width: 5, height: 10},
		&Circle{radius: 7},
		&Triangle{side: 3},
	}

	results := make(chan Result, len(shapes))
	var wg sync.WaitGroup

	// สร้าง Goroutines สำหรับ Worker
	for i, shape := range shapes {
		wg.Add(1)
		go func(i int, s Shape) {
			defer wg.Done()
			area := s.Area()
			perimeter := s.Perimeter()
			results <- Result{Area: area, Perimeter: perimeter}
			fmt.Printf("Worker %d: Area = %.2f, Perimeter = %.2f\n", i, area, perimeter)
		}(i, shape)
	}

	wg.Wait()
	close(results)

	// อ่านผลลัพธ์จาก Channel
	var totalArea, totalPerimeter float32
	for result := range results {
		totalArea += result.Area
		totalPerimeter += result.Perimeter
	}

	fmt.Printf("Total Area: %.2f\n", totalArea)
	fmt.Printf("Total Perimeter: %.2f\n", totalPerimeter)
}

// Result struct สำหรับเก็บผลลัพธ์
type Result struct {
	Area      float32
	Perimeter float32
}

// Shape interface
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
	radius float32
}

func (c *Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) Perimeter() float32 {
	return 2 * math.Pi * c.radius
}

// Triangle
type Triangle struct {
	side float32
}

func (t *Triangle) Area() float32 {
	return float32(math.Sqrt(3)) / 4 * t.side * t.side
}

func (t *Triangle) Perimeter() float32 {
	return 3 * t.side
}
