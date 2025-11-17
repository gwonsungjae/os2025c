package main

import (
	"fmt"

	"github.com/headfirstgo/magazine"
)

func main() {
	var s1 magazine.Subscriber
	var e1 magazine.Employee
	s1.Name = "Choi Inha"
	e1.Name = "Lee Inha"
	e1.Salary = 50000000
	e1.Address.City = "Incheon"
	s1.Address.City = "Seoul"
	fmt.Println(s1.Name, s1.Address.City)
	fmt.Println(e1.Name, e1.Salary)
}
