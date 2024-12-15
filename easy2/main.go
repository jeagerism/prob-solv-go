// เขียนฟังก์ชันที่รับสตริงมา 1 ตัว และตรวจสอบว่าสตริงนั้นเป็น Palindrome หรือไม่ (คำที่อ่านจากซ้ายไปขวาและขวาไปซ้ายเหมือนกัน)
// ตัวอย่าง Input:
// "madam"
// Output ที่คาดหวัง:
// true

// ตัวอย่าง Input:
// "hello"
// Output ที่คาดหวัง:
// false

// Tips:
// ใช้ตัวชี้ตำแหน่ง (pointers) ซ้าย-ขวาเพื่อเปรียบเทียบตัวอักษรจากสองด้าน
// หรือลองใช้ reverse string แล้วเปรียบเทียบกับต้นฉบับ

package main

import "fmt"

func main() {
	input := "madam"
	if input == reverse(input) {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
	fmt.Println(reverse(input))
}

func reverse(str string) string {
	var result string
	for _, str := range str {
		result = string(str) + result
	}
	return result
}
