package main

import "fmt"

func main() {
	a := 43
	var b *int = &a

	fmt.Println(a)
	fmt.Println(b)

	*b = 42
	fmt.Println(a)

}
