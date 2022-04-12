package main

import "fmt"

var packVar = "Package Scope"
var funcVar = "Func Scope"

//packVar := "Package Scope" Bu şekilde hata verir

func main() {
	//var funcVar = "Func Scope"
	//funcVar := "Func Scope"
	fmt.Println(funcVar)
	fmt.Println(packVar)

}
