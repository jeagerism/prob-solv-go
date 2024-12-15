package main

import (
	"errors"
	"fmt"
)

func main() {
	var p1, p2 int
	fmt.Print("Enter number p1 p2 : ")
	fmt.Scan(&p1, &p2)

	result, err := DivideNumbers(p1, p2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Divided = %d", result)
	}
}

func DivideNumbers(p1, p2 int) (int, error) {
	// ตรวจสอบว่าตัวหารเป็น 0 หรือไม่
	if p2 == 0 {
		return 0, errors.New("Cannot divide by zero")
	}
	d := p1 / p2
	return d, nil
}
