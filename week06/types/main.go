package main

import (
	"fmt"
	"reflect"
)

func main() {
	//fmt.Println(math.Round(2.21))
	//fmt.Println(math.Ceil(2.21))
	//fmt.Println(strings.Title("go developer!"))
	//fmt.Println("Kim\nInha\t\"\\")
	//fmt.Println('2', '가')
	fmt.Println(reflect.TypeOf(2.31))
	fmt.Println(reflect.TypeOf("sungjae"))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf('A'))
	fmt.Println(reflect.TypeOf(19))
}
