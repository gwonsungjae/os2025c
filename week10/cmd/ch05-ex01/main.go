package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var arrayBool [3]bool = [3]bool{true, false, true}
	arrayBool := [3]bool{true, false, true}
	var arrayInt [3]int
	fmt.Println(reflect.TypeOf(arrayBool))
	fmt.Println("%#v\n", arrayBool)

	fmt.Println(arrayBool[1])
	arrayInt[1] = 2
	fmt.Println(arrayInt[1])

	//fmt.Println((arrayBool[1]))
	//arrayInt[1]++
	//arrayInt[1] = arrayInt[1] + 1
	//fmt.Println(arrayInt[1])
}
