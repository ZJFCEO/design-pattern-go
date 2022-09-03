package main

import (
	"fmt"
)

func main() {
	api := NewAPI("zh")
	fmt.Println(api.Say("泡泡"))

	api2 := NewAPI("en")
	fmt.Println(api2.Say("bubble"))
}
