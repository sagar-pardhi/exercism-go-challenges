package main

import (
	"fmt"
	"time"
)

func main() {
	date := "6/6/2005 10:30:00"
	layout := "2/1/2006 15:04:05"
	t, _ := time.Parse(layout, date)

	fmt.Println(t.Format("Monday, January 2, 2006, at 15:04"))

	anniversaryDate := time.Date(2020, time.September, 15, 0, 0, 0, 0, time.UTC)
	fmt.Println(anniversaryDate)
}
