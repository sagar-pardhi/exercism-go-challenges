package main

import (
	"fmt"
)

func main() {
	var float = -12.345

	res := fmt.Sprintf("This is the number %.1f", float)
	fmt.Println(res)
}
