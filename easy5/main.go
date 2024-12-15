// ผลรวมของตัวเลขที่ไม่ซ้ำกัน
// คำอธิบาย:
// เขียนฟังก์ชันที่รับอาเรย์ของตัวเลข และคืนค่าผลรวมของตัวเลขที่ไม่ซ้ำกันในอาเรย์

// ตัวอย่าง Input:
// []int{1, 2, 2, 3, 4, 4, 5}
// Output ที่คาดหวัง:
// 9 // เพราะ 1 + 3 + 5

// Tips:
// ใช้ map เพื่อตรวจสอบว่าตัวเลขเคยถูกเพิ่มหรือไม่
// ใช้ for loop เพื่อตรวจสอบซ้ำ

package main

import "fmt"

func main() {
	input := []int{1, 2, 2, 3, 4, 4, 5}
	sum := 0
	m := make(map[int]int)
	for _, v := range input {
		m[v]++
	}

	for key, count := range m {
		if count == 1 {
			sum += key
		}
	}
	fmt.Println(sum)
}
