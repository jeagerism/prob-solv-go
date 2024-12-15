// เขียนฟังก์ชันที่รับอาเรย์ของตัวเลขจำนวนเต็ม และคืนค่าอาเรย์ที่เรียงลำดับจากน้อยไปมาก

// ตัวอย่าง Input:
// []int{4, 2, 7, 1, 3}
// Output ที่คาดหวัง:
// []int{1, 2, 3, 4, 7}

// Tips:
// ใช้ sort.Ints() จากแพ็กเกจ sort ใน Go
// หรือเขียนอัลกอริธึมการเรียงลำดับด้วยตนเอง เช่น Bubble Sort (สำหรับฝึก)

package main

import "fmt"

func main() {
	input := []int{4, 2, 7, 1, 3}
	smalest(input)
	largest(input)
}

func largest(arr []int) {
	n := len(arr)
	fmt.Println("=== มากไปน้อย ===")
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				fmt.Println(arr)
			}
		}
	}
}

func smalest(input []int) {
	n := len(input)
	fmt.Println("=== น้อยไปมาก ===")
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
				fmt.Println(input)
			}
		}
	}
}
