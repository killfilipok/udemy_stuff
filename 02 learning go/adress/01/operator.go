package main

import "fmt"

func main() {
	a := 42

	fmt.Println("memory adress of a is: ", &a)
	fmt.Printf("%d ", &a)
}
