package main

import "fmt"

func main() {
	ch := make(chan int) // int 채널 생성

	go func() {
		ch <- 123 // 채널에 값 보내기
	}()

	x := <-ch // 채널에서 값 받기
	fmt.Println(x)
}
