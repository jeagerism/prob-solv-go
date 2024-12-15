package main

import "fmt"

func main() {
	// ตัวเลขที่ใช้คำนวณ
	numbers := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	// แบ่ง slice เป็น 2 ส่วน
	mid := len(numbers) / 2
	part1 := numbers[:mid] // {10, 20, 30, 40, 50}
	part2 := numbers[mid:] // {60, 70, 80, 90, 100}

	// สร้าง channel เพื่อรับผลลัพธ์จาก goroutines
	results := make(chan int, 2)

	// ใช้ goroutines คำนวณผลรวม
	go sum(part1, results)
	go mul(part2, results)

	// รับผลลัพธ์จาก goroutine
	result1 := <-results
	result2 := <-results

	// รวมผลลัพธ์ทั้งหมด
	total := result1 + result2

	fmt.Printf("Total sum: %d\n", total)
}

// ฟังก์ชันคำนวณผลรวม
func sum(num []int, results chan int) {
	sum := 0
	for _, v := range num {
		sum += v
	}
	results <- sum
}

// ฟังก์ชันคำนวณผลคูณ
func mul(num []int, result chan int) {
	product := 1 // เริ่มต้นเป็น 1 เพื่อคำนวณการคูณ
	for _, v := range num {
		product *= v
	}
	result <- product
}
