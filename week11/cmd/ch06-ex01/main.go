// Package datafile allows reading data samples from files.
package main

import "fmt"

func main() {
	//var sub []string
	//sub = make([]string, 3)
	//sub := make([]string, 3)
	sub := []string{"Go", "Javascript", "Python", "Linux"}
	//sub[0] = "Go"
	//sub[2] = "Python"

	subSlice := sub[1:3]
	fmt.Println("===============")
	for i := 0; i < len(subSlice); i++ {
		fmt.Println(subSlice[i])
	}

	for _, subject := range sub {
		fmt.Println(subject)
	}
}
