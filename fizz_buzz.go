package main 
import "fmt"

func main() {
  
// fizz, buzz - numbers between 1 and 20
// if num/3 = 0 - fizz
// if num/5 = 0 - buzz
// if num/ && /5 - fizzbuzz

for i := 1; i < 21; i++ {

  if i % 15 == 0 {  //can be i % 3 == 0 && i % 5 == 0
    fmt.Println("fizzbuzz")
  } else if i % 3 == 0 {
    fmt.Println("fizz")
  } else if i % 5 == 0 {
    fmt.Println("buzz")
  } else {
    fmt.Println(i)
  }
  }
}

