package main

import "fmt"

func main() {
	var votes int
	votes = 3

	var voteCounter *int
	voteCounter = &votes

	if voteCounter != nil {
		fmt.Println(*voteCounter)
	} else {
		fmt.Println(0)
	}

}
