package main 
import "fmt"

func main() {
  //Strings in Go are IMMUTABLE - can't change - book[0] = "n" will throw an error

  book := "The color of magic"
  fmt.Println(book) // The color of magic

  fmt.Println(len(book)) // 18

  fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])
  // book[0] = 84 (type uint8)
  // uint8 is a byte

  //Slicing
  //Slice (start, end), 0 based, 1/2 empty range
  fmt.Println(book[4:11]) // color o

  // Slice (no end)
  fmt.Println(book[4:]) // color of magic

  //Slice (no start)
  fmt.Println(book[:5]) // The c

  //Concatenate
  fmt.Println(book + " is ...") // The color of magic is ...

  //The strings are in unicode so we can use special characters

  //Multi line
  poem := `
  The road goes...
  Down ...
  ...
  `

  fmt.Println(poem) // will print the multiline string as it is
}
