package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("1st")
	defer fmt.Println("2nd")
	defer fmt.Println("3rd")
	fmt.Println("main")
}

// func main() {
// 	file, err := os.Open("test.txt")
// 	if err != nil {
// 		fmt.Println("파일 오픈 실패:", err)
// 		return
// 	}

// 	// defer로 파일 닫기를 미리 예약
// 	defer file.Close()

// 	// 파일 작업 수행
// 	fmt.Println("파일 읽는 중...")
//}
