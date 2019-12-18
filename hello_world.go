//we have one line and multiline (/* */) comments 

package main 
//go is organized in packages - helps organizing the code and with big projects
//main is a special package in go - it helps with the code compilation

import "fmt"
//go is a simple language without many libraries included - most of the libraries will be imported in addition

//the main function will be executed by the go runtime when the program starts
//go supports unicode by default 
func main() {
  fmt.Println("Hello World")
}
