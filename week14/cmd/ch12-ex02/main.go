package main

import "fmt"

func safeDivide(a, b int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("에러 발생:", err)
		}
	}()

	if b == 0 {
		panic("0으로 나눌 수 없습니다!")
	}

	result := a / b
	fmt.Println("결과:", result)
}

func main() {
	fmt.Println("첫 번째 호출")
	safeDivide(10, 2)

	fmt.Println("\n두 번째 호출")
	safeDivide(10, 0)

	fmt.Println("\n프로그램 계속 실행됨")
}
