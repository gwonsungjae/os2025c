package main

import (
	"fmt"
)

func main() {
	//var arrayBool [3]bool = [3]bool{true, false, true}
	arrayBool := [2]bool{true, false}
	//var arrayInt [3]int
	arrayInt := [3]int{-9, 11, 7}
	for i := 0; i < len(arrayBool); i++ {
		fmt.Println(i, arrayBool[i])
		fmt.Println(i, arrayInt[i])
	}

	//fmt.Println(reflect.TypeOf(arrayBool))
	//fmt.Println("%#v\n", arrayBool)

	fmt.Println(arrayBool[1])
	arrayInt[1] = 2
	fmt.Println(arrayInt[1])

	//fmt.Println((arrayBool[1]))
	//arrayInt[1]++
	//arrayInt[1] = arrayInt[1] + 1
	//fmt.Println(arrayInt[1])
}
