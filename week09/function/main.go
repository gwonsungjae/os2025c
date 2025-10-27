package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//	func swap(first *int, second *int) {
//		var temp int = 0
//		temp = *first
//		*first = *second
//		*second = temp
//		fmt.Println(first, second)
//	}
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}
func main() {

	// a, b := 10, 20
	// fmt.Println(a, b)
	// swap(&a, &b)
	// fmt.Println(a, b)
	//fmt.Printf("%0.3f\n", math.Sqrt(-9.3))

	fmt.Print("점수 입력 : ")
	score, err := GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	passFail := ""
	if score >= 60 {
		passFail = "합격!"
	} else {
		passFail = "불합격!"
	}
	fmt.Print("%.2f점은 %v\n", score, passFail)
}
