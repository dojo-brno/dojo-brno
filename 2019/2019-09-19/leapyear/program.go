package main

import "fmt"
import "leapyear"

func main() {
  var rok int
  var p string
  fmt.Printf("Zadej rok:")
  _,chyba := fmt.Scanf("%v",  &rok)
  if chyba != nil {
    fmt.Printf("Spatne zadany rok.\n")
    return
  }
  p = "neni"
  if leapyear.LeapYear(rok) {
  p = "je"
  }
  
  fmt.Printf("Zadany rok %v %v prestupny.\n", rok,p)
}