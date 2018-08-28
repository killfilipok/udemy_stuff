package main

import "fmt"

func main() {

	for i := 1; i < 200; i++ {
		if i%2 == 1 {
			fmt.Println("ODD")
		} else {
			fmt.Println("EVEN")
		}
	}
}
