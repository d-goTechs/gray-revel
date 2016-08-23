package main

import "fmt"

func main() {


	var i int= 5
	switch i {
	case 1: fmt.Print("1")
	case 2: fmt.Print("2")
	case 3: fmt.Print("3")
	case 4,5,6: fmt.Print("bigger than 3")
	default:
	fmt.Print("default")
	}
}
