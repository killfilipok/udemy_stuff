package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string
	Age     int
	Premium bool
}

func main() {
	var nillable Person

	fmt.Println(nillable)

	fakeJson := []byte(`{"Name": "Fackes","Age" : 18,"Premium":true}`)

	json.Unmarshal(fakeJson, &nillable)

	fmt.Println(nillable)
}
