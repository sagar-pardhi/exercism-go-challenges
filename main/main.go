package main

import (
	"fmt"
)

func main() {
	var balance float64 = 200.75

	switch {
	case balance < 0:
		fmt.Println(float32(3.213))
	case balance >= 0 && balance < 1000:
		fmt.Println(float32(0.5))
	case balance >= 1000 && balance < 5000:
		fmt.Println(float32(1.621))
	default:
		fmt.Println(float32(2.475))
	}

	fmt.Println("Interest Rate", float64(balance*0.5*1/100))

	count := 0
	for balance < 210 {
		balance += 0.5 + balance
		count++
	}

	fmt.Println(count, balance)
}
