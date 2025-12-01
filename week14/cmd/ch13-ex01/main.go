package main

import (
	"fmt"
	"time"
)

func say(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg, ":", i)
		time.Sleep(2 * time.Second)
	}
}

func hi(msg string) {
	time.Sleep(7 * time.Second)
	fmt.Println("안녕", msg)
}

func main() {
	start := time.Now()
	go hi("고루틴1")               // 새 고루틴에서 실행
	go say("고루틴2")              // 새 고루틴에서 실행
	time.Sleep(8 * time.Second) // 위 고루틴들을 대기
	fmt.Println("전체 실행 시간 : ", time.Since(start))
}
