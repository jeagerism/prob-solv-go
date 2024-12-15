package main

import "fmt"

func main() {
	arr1 := []int{5, 10, 15}
	n := len(arr1)
	fmt.Println("AVG Array 1 : ", avg(n, arr1))
	//--//
	var newNum int
	fmt.Print("Enter a number to add: ")
	fmt.Scanln(&newNum)
	arr2 := append(arr1, newNum)
	fmt.Println("AVG Array 2 : ", avg(len(arr2), arr2))
}

func avg(len int, arr []int) float32 {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	avr := float32(sum) / float32(len)
	return avr
}
