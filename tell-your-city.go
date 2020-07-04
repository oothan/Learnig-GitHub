package main

import (
  "fmt"
  "os"
  "bufio"
)

func main() {
  reader := bufio.NewReader(os.Sdtin)
  fmt.Print("Tell Your City : ")
  city, _ := reader.ReadString('\n')
  fmt.Println("Your city is ", city)
}
