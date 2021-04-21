package main

import ("fmt")

func main() {
  var name string
  
  fmt.Println("What's your name? <> ")
  fmt.Scanln(&name)
  if name == "ziksthemw" {
    fmt.Println("Hello my king!")
  } else {
    fmt.Println("Hello, " + name + "!")
  }
}
