package main 
import "fmt"

func main() {

  // var x int //if we have 2 int the division (/) will return int
  // if we have 1 int and 1 float, it will return an error - go is strongly typed
  // that's why the types need to match - we need 2 floats to have the correct result
  /* Declare then assign:

  var x float64
  var y float64

  x = 1
  y = 2 
  
  */

  //Declare + assign:
  
  x := 1
  y := 2
  //can be in 1 line as well: x, y := 1, 2

  fmt.Printf("x=%v, type of %T\n", x, x)
  fmt.Printf("y=%v, type of %T\n", y, y)

  /* var mean float64 // the ruslt can't be int - needs to be float as well
  mean = (x + y) / 2 */

  mean := (x + y) / 2
  // here we are using an integer division, which returns int, e.g. result = 1
  fmt.Printf("result: %v, type of %T\n", mean, mean) // 1.5
  
}
