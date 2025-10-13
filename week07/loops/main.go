package main

import (
	"fmt"
	"reflect"
)

func main() {
	var length float64 = 3.2
	var width int = 2
	fmt.Println("면적은", int(length)*width)
	fmt.Println("length > width?", int(length) > width)
	fmt.Println("형변환", reflect.TypeOf(int(length)))
	fmt.Println("원본", reflect.TypeOf(length))
}
