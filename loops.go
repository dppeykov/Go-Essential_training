package main 
import "fmt"

func main() {
  
  // for loop
  for i := 0; i <3; i++ {
      fmt.Println(i) //0,1,2
  }

  fmt.Println("----")
  for i := 0; i < 3; i++ {
      if i > 1 {
          break
      }
      fmt.Println(i) //0,1
  }

  fmt.Println("----")
  for i := 0; i < 3; i++ {
      if i < 1 {
          continue
      }
      fmt.Println(i) //1,2
  }

  //while loop

  fmt.Println("----")
  a := 0
  for a < 3 {
      fmt.Println(a) //0,1,2
      a++
  }

  fmt.Println("----")
  b := 0
  for { //for without any conditions is like while true loop in the other languages
      if b > 2 {
          break
      }
      fmt.Println(b) //0,1,2
      b++
  }

}
