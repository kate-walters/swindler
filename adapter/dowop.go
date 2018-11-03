package adapter

import (
	"fmt"
)

type Dowop struct {}

func (a Dowop) Add(thing string) {
	fmt.Print("Enter hostname: ")
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
}
