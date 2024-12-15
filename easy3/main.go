// เขียนฟังก์ชันเพื่อหาจำนวนเฉพาะในช่วงตัวเลขตั้งแต่ 2 ถึง n

// ตัวอย่าง Input:
// 10
// Output ที่คาดหวัง:
// []int{2, 3, 5, 7}

// Tips:
// ใช้การวนซ้ำ for เพื่อตรวจสอบแต่ละตัวเลข
// ใช้ฟังก์ชันเสริมเพื่อตรวจสอบว่าเป็นจำนวนเฉพาะหรือไม่
// จำนวนเฉพาะต้องมากกว่า 1 และหารด้วยตัวเลขอื่นนอกจาก 1 และตัวมันเองไม่ได้

package main

import "fmt"

func main() {
	input := 100
	output := []int{}
	for i := 2; i <= input; i++ {
		if isPrime(i) {
			output = append(output, i)
		}
	}
	fmt.Println(output)
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
