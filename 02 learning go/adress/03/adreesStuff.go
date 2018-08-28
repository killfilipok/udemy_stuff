package main

import "fmt"

func main() {
	a := 42
	var b *int = &a
	fmt.Println(b)
	fmt.Println(*b)
}
