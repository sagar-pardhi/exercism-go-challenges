package main

import (
	"fmt"
)

func main() {
	birdsPerDay := []int{2, 5, 0, 7, 4, 1}

	for i := 0; i < len(birdsPerDay); i += 2 {
		fmt.Println(birdsPerDay[i])
		birdsPerDay[i] += 1
	}

	fmt.Println(birdsPerDay)

	// birdCount := 0
	// for i := 0; i < len(birdsPerDay); i++ {
	// 	birdCount += birdsPerDay[i]
	// }

	// fmt.Println(birdCount)
}
