package main

import "fmt"

type orderStatus int
func changeStatus(status orderStatus)  {
	fmt.Println( "Status changed to", status)
}
const (
	received orderStatus = iota
	  confirmed
	  prepared
	  shipped
	  delivered
  )
func main() {
	
	
	changeStatus(received)
}
