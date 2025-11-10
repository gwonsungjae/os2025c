// Package datafile allows reading data samples from files.
package main

import "fmt"

func main() {
	//var sub []string
	//sub = make([]string, 3)
	//sub := make([]string, 3)
	sub := [4]string{"Go", "Javascript", "Python", "Linux"}
	//sub[0] = "Go"
	//sub[2] = "Python"

	subSlice := sub[:3]
	//sub[0] = "Java"
	subSlice[0] = "Database"
	for _, subject := range sub {
		fmt.Println(subject)
	}

	fmt.Println("===============")
	for i := 0; i < len(subSlice); i++ {
		fmt.Println(subSlice[i])
	}

}
