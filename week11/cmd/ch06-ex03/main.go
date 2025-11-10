// average2 calculates the average of several numbers.
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func mean(numbers ...float64) float64 {
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

func main() {
	arguments := os.Args[1:]
	var numbers []float64
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	fmt.Printf("평균 고기 소비량: %0.2f\n", mean(numbers...))
}
