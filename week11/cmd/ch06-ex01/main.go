// Package datafile allows reading data samples from files.
package main

import "fmt"

func main() {
	//var sub []string
	//sub = make([]string, 3)
	sub := make([]string, 3)
	sub[0] = "Go"
	sub[2] = "Python"

	for _, subject := range sub {
		fmt.Println(subject)
	}
}
