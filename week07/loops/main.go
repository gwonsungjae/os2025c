package main

import (
	"fmt"
	"strings"
	"strconv"
	"bufio"
)

func main() {
	fmt.Print("Enter a score: ")
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	i = strings.TrimSpace(i)                // 문자열 주위에 붙은 공란 및 탭 키 등 제거
	score, err := strconv.ParseFloat(i, 64) // 정리된 문자열을 실수타입으로 변환
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if score >= 60 {
		status = "Pass"
	} else {
		status = "Fail"
	}
	fmt.Println(score, status)}


	//var float64 float64 = 2.71
 	//var f float64 = 3.001
	//fmt.Println(float64)


	//var fmt float64 = 2.71
	//r := bufio.NewReader(os.Stdin)
	//i, _ := r.ReadString('\n') /// ignore error
	//i, err := r.ReadString('\n')
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(i)

	//var now time.Time = time.Now()

	//var month time.Month = now.Month() // month := now.Month()
	//var day int = now.Day()
	//fmt.Println(month, day)

	// var univ string = "Go$ Inha$"
	//changer := strings.NewReplacer("$", "!")
	// changed := changer.Replace(univ)
	//changed := changer.Replace("Go$ Inha$")
	//fmt.Println(changed)

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
