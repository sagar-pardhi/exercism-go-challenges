package main

import (
	"fmt"
)

type Resident struct {
	name    string
	age     int
	address map[string]string
}

func main() {
	addr := map[string]string{}

	r := Resident{
		name:    "",
		age:     10,
		address: addr,
	}

	fmt.Println(len(r.address))
}
