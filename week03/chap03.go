package main

import "fmt"

func main() {
	// a := 4
	// pa := &a
	var a int = 4
	var pa *int = &a

	fmt.Println(*pa)
	fmt.Println(pa)
	fmt.Println(&a)
}
