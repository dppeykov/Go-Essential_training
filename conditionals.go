package main 
import "fmt"

func main() {

  x := 10

  if x > 5 {
      fmt.Println("x is big")
  } else {
      fmt.Println("x is not that big")
  }

  if x > 5 && x < 15 {
      fmt.Println("x is just right")
  }

    if x < 20 || x > 30 {
      fmt.Println("x is out of range")
  }

  a := 11.0
  b := 20.0

  if frac := a / b; frac > 0.5 {
      fmt.Println("a is more than half of b")
  }

// swith statement:
 
  x_switch := 5

  switch x_switch {
  case 1:
    fmt.Println("one")
  case 2:
    fmt.Println("two")
  case 3:
    fmt.Println("two")
  default:
    fmt.Println("many")
  }

}
