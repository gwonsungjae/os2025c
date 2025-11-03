package main

import (
	"fmt"
	"log"
	"week10/pkg/keyboard"
)

func main() {
	fmt.Print("화씨 온도 입력 : ")
	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	var celsius float64
	celsius = (fahrenheit - 32) * 5 / 9
	fmt.Printf("화씨 %.2f도는 섭씨 %.2f도 입니다.", fahrenheit, celsius)
}
