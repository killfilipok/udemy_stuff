package main

import "fmt"

type myType struct {
	greeting string
	message  string
}

func main() {
	var x myType
	x.greeting = ""
	x.message = "sd"
	checkType(x)
}

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case myType:
		fmt.Println("myType")
	default:
		fmt.Println("default")
	}
	if y := "lol"; y == x {
		fmt.Println(`string as "lol"`, y)
	}
}
