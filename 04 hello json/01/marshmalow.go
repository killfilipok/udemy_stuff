package main

import "fmt"
import "encoding/json"

type User struct {
	Name     string
	Email    string
	Password int  `json:"-"`
	Premium  bool `json:"subscribed?"`
}

func main() {
	philip := User{"philip", "killfilipok@gmail.com", 123321, true}

	jsonObj, _ := json.Marshal(philip)

	fmt.Println(jsonObj)
	fmt.Println(string(jsonObj))
}
