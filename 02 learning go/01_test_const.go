package main

import "fmt"

const (
	_  = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
	GB = 1 << (iota * 10)
	TB = 1 << (iota * 10)
)

func main() {
	fmt.Printf("%d \t %b \n ", KB, KB)
	fmt.Printf("%d \t %b \n ", MB, MB)
	fmt.Printf("%d \t %b \n ", GB, GB)
	fmt.Printf("%d \t %b \n ", TB, TB)
}
