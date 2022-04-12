package main

import "fmt"

var packVar = "Package Scope"

func main() {
	var funcVar = "Func Scope"
	/*if true {
		var blockVar = "Block Scope"
		fmt.Println(blockVar)
	}
	fmt.Println(blockVar) //Burda scope dışında olduğu için hata verir*/

	fmt.Println(funcVar)
	fmt.Println(packVar)

	myFunc()

}
func myFunc() {
	fmt.Println(packVar) //Hata vermez
	//fmt.Println(funcVar) Hata verir
}
