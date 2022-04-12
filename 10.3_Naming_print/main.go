package main

import "fmt"

var x = 10

func main() {

	fmt.Println(x)
	myFunc()
}
func myFunc() {
	fmt.Println(x)
}
