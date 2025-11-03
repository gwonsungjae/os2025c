package main

import "fmt"

func main() {
	var arrayBool [3]bool
	var arrayInt [3]int
	fmt.Println((arrayBool[1]))
	arrayInt[1]++
	arrayInt[1] = arrayInt[1] + 1
	fmt.Println(arrayInt[1])
}
