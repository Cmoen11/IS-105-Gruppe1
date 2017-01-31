package main

import "os"
import "fmt"

func main(){
  argsWithProg := os.Args
  argsWithoutProg := os.Args[8:]

  arg := os.Args[2]

  fmt.Println(argsWithProg)
  fmt.Println(argsWithoutProg)
  fmt.Println(arg)
}
