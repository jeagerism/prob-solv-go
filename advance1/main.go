package main

import (
	"fmt"
	"math"
)

// โจทย์ 1: ระบบจัดการรูปทรงด้วย Slice และการคำนวณรวม
// คำอธิบาย:
// สร้าง ระบบจัดเก็บรูปทรง ใน Slice ([]Shape) โดยสามารถเพิ่มรูปทรงใหม่ได้ เช่น Rectangle, Circle, Triangle
// สร้างฟังก์ชันเพื่อ:
// คำนวณ พื้นที่รวม (Total Area) ของรูปทรงทั้งหมด
// คำนวณ เส้นรอบรูปทั้งหมด (Total Perimeter)
// แสดงผลลัพธ์เป็นแบบ เรียงตามลำดับที่เพิ่มเข้าไป
// ไอเดียในการเริ่มต้น:
// ใช้ Slice เก็บรูปทรง (เช่น shapes := []Shape{})
// ใช้ Polymorphism และ Interface ในการเรียก Area และ Perimeter

// output =>
// Total Area: 50 + 153.94 + 3.897 ≈ 207.837
// Total Perimeter: 30 + 43.98 + 9 ≈ 82.98

func main() {
	r := &Rectangle{width: 5, height: 10}
	c := &Circle{radius: 7}
	t := &Triangle{side: 3}

	shape := []Shape{}
	shape = append(shape, r)
	shape = append(shape, c)
	shape = append(shape, t)

	var totalArea float32
	var totalPerimeter float32

	for _, v := range shape {
		totalArea += v.Area()
		totalPerimeter += v.Perimeter()
	}

	fmt.Printf("Total Area: %.2f + %.2f + %.2f ≈ %.2f\n", r.Area(), c.Area(), t.Area(), totalArea)
	fmt.Printf("Total Perimeter: %.2f + %.2f + %.2f ≈ %.2f", r.Perimeter(), c.Perimeter(), t.Perimeter(), totalPerimeter)
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
