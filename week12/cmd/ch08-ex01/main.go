package main

import "fmt"

type subscriber struct {
	name  string
	price int
}

func applyPrice(s *subscriber) {
	s.price = 10000
	s.name = "Kim Inha"
}
func main() {
	var s1 subscriber
	//s1.name = "Lee Inha"
	applyPrice(&s1)
	fmt.Println(s1.name, s1.price)
}
