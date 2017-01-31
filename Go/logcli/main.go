package main

import (
  "os"
  "fmt"
  "./log"
)

func main() {
  argsWithProg :=os.Args
  argsWithoutProg := os.Args[1:]

  arg := os.Args[3]
   fmt.Println(argsWithProg)
   fmt.Println(argsWithoutProg)
   fmt.Println(arg)
  
  fmt.Println("utregning fullført")
}
