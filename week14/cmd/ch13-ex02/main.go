package main

import "fmt"

func abc(channel chan string) {
	channel <- "내고향 스페셜\n"
	channel <- "KBS 뉴스광장\n"
	channel <- "인간극장\n"
}

func def(channel chan string) {
	channel <- "건강의 재구성 썰록(재)\n"
	channel <- "오늘N\n"
	channel <- "찾아가는 꾸러기 교실\n"
}

func main() {
	kbs := make(chan string)
	mbc := make(chan string)
	go abc(kbs)
	go def(mbc)
	fmt.Print(<-kbs)
	fmt.Print(<-mbc)
	fmt.Print(<-kbs)
	fmt.Print(<-mbc)
	fmt.Print(<-kbs)
	fmt.Print(<-mbc)
	fmt.Println()
}
