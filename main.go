package main

import "fmt"

func main() {
  fmt.Println("Hello World and welcome to the ID compiler (GO version)")
  fmt.Println("Type your name:")
  fmt.Scanf("%s", &name)
  fmt.Println("OK. Type your age and date of birth:")
  fmt.Scanf("%s", &dob)
  fmt.Println("OK. Type your nationality:")
  fmt.Scanf("%s", &nationality)
  fmt.Println("OK. Type your parents name:")
  fmt.Scanf("%s", &parents)
  fmt.Println("Alright.Your info is: ", &name, &dob, &nationality, &parents)
}
