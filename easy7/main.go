// หาความยาวของคำที่ยาวที่สุดในสตริง
// คำอธิบาย:
// เขียนฟังก์ชันที่รับสตริงเป็นประโยค และคืนค่าความยาวของคำที่ยาวที่สุดในประโยคนั้น

// ตัวอย่าง Input:
// "This is a wonderful example"
// Output ที่คาดหวัง:
// 9 // "wonderful"

// Tips:
// แยกคำในสตริงด้วย strings.Fields()
// ใช้ len() เพื่อหาความยาวของคำแต่ละคำ

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "This is a wonderful example"
	arr := strings.Fields(str)

	n := len(arr)
	l := largest{len: len(arr[0]), str: ""}

	for i := 0; i < n; i++ {
		if len(arr[i]) > l.len {
			l.len = len(arr[i])
			l.str = arr[i]

		}
	}
	fmt.Println(l)
}

type largest struct {
	len int
	str string
}
