package main

import "fmt"

const mToYards float64 = 1.09361

func main() {
	var meters float64

	fmt.Scan(&meters)

	yards := meters * mToYards

	fmt.Println(yards)
}
