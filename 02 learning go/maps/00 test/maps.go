package main

import "fmt"

func main() {
	myMap := map[string]int{
		"lol":      0,
		"wow":      1,
		"huh":      2,
		"dissmiss": 3,
	}

	var x bool
	var y int
	y, x = myMap["lol"]

	for k, v := range myMap {
		fmt.Println(k, "--", v)
	}

	if val, exist := myMap["lol"]; exist {
		delete(myMap, "lol")
		fmt.Println("val was: ", val)
		fmt.Println("now :", myMap["lol"])
	}
	fmt.Println(myMap)
	fmt.Println(x, y)
}
