package main

import (
  "fmt"
)

func main() {
  var name string
  var age int

  fmt.Println("What is your name, friend?")
  fmt.Scan(&name)
  fmt.Println("Hello", name + "!")

  fmt.Println("How old are you?")
  fmt.Scan(&age)
  fmt.Printf("So you're %d years old? Nice to meet you, %v!", age, name)
  fmt.Println()

  newMessage := fmt.Sprintf("To think, that you have been on this planet for %d years. Imagine that, %v!", age, name)
  fmt.Println(newMessage)
}
