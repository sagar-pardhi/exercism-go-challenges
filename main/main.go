package main

import (
	"fmt"
)

func main() {
	friendsList := []string{"sauce", "noodles", "béchamel", "marjoram"}
	myList := []string{"sauce", "noodles", "meat", "tomatoes", "?"}

	lastElem := friendsList[len(friendsList)-1]

	myList[len(myList)-1] = lastElem

	fmt.Println(myList)
}
