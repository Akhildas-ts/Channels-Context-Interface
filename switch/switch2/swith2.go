package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("when's staurday")
	today := time.Now().Weekday()
	switch time.Saturday {

	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("tommarow")
	case today + 2:
		fmt.Println("in 2 days ")

	default:
		fmt.Println("too far away")
	}
}
