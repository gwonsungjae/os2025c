package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var now time.Time = time.Now()
	var month time.Month = now.Month() // month := now.Month()
	var day int = now.Day()
	fmt.Println(month, day)

	// var univ string = "Go$ Inha$"
	changer := strings.NewReplacer("$", "!")
	// changed := changer.Replace(univ)
	changed := changer.Replace("Go$ Inha$")
	fmt.Println(changed)

	//var now time.Time = time.Now()
	//month := now.Month()
	//var day int = now.Day()
	//fmt.Println(month, day)

	//var length float64 = 3.2
	//var width int = 2
	//fmt.Println("면적은", int(length)*width)
	//fmt.Println("length > width?", int(length) > width)
	//fmt.Println("형변환", reflect.TypeOf(int(length)))
	//fmt.Println("원본", reflect.TypeOf(length))
}
