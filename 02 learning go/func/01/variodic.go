package main

import "fmt"

func main()  {
  foo(1,85,16,864,16)
}

func foo(arg ...float64)  {
  var temp float64
  for _, v:= range arg  {
    if v > temp {
      temp = v
    }
  }
  fmt.Println(temp)
}
