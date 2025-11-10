package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetFloats(fileName string) ([3]float64, error) {
	var numbers [3]float64
	file, err := os.Open(fileName)
	if err != nil {
		return numbers, err
	}
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		i++
	}
	err = file.Close()
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, nil
}
func main() {
	weights, err := GetFloats("meatWeight.txt")
	if err != nil {
		log.Fatal(err)
	}
	sum := 0.0
	for i := 0; i < len(weights); i++ {
		sum = sum + weights[i]
	}
	weeks := float64(len(weights))
	fmt.Println("평균 고기 소비량 : ", sum/weeks)
}
