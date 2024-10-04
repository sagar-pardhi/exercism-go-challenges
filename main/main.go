package main

import (
	"fmt"
)

func main() {
	units := map[string]int{}
	units["quarter_of_a_dozen"] = 3
	units["half_of_a_dozen"] = 6
	units["dozen"] = 12
	units["small_gross"] = 120
	units["gross"] = 144
	units["great_gross"] = 1728

	bill := map[string]int{}

	bill["carrot"] += units["dozen"]
	bill["carrot"] += units["half_of_a_dozen"]

	// bill["carrot"] -= units["quarter_of_a_dozen"]

	value, exists := bill["carrot"]

	fmt.Println(value, exists)
	fmt.Println(bill)
}
