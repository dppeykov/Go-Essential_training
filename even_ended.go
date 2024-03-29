package main 
import "fmt"

func main() {
  //count the number of even ended numbers
  //count = 0
  count := 0

  // for every pair of 4 digit numbers
  for a := 1000; a <= 9999; a++ {
      for b := a; b <= 9999; b++ { // don't count twice
        n := a * b

        // if a * b even ended
        s := fmt.Sprintf("%d", n)
        if s[0] == s[len(s)-1] {
            //count = count + 1
            count++
        }
      }
  }
  //pring count
  fmt.Println(count)
}
