package adapter

import (
	"fmt"
)

type Noop struct {}

func (a Noop) Add(thing string) {
	fmt.Print("Enter name of file: ")
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
}
