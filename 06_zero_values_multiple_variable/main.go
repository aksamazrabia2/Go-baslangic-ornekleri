package main

import "fmt"

func main() {

	/*	var (
		name      string  = "Rabia"
		age       int     = 40
		isMarried bool    = true
		weight    float32 = 72.5
		height    int     = 172
	)*/
	/*var name, age, isMarried, weight, height = "Rabia", 40, true, 72.5, 172*/
	/*name, age, isMarried, weight, height := "Rabia", 40, true, 72.5, 172*/

	var name string
	fmt.Println(name) //boş string değeri gösterir.Zero values
	var age int
	fmt.Println(age) //sayısal ifadeler içinse çalışınca sıfır sonucunu verir.numerik
	var height float32
	fmt.Println(height) //sayısal ifadeler içinse çalışınca sıfır sonucunu verir.numerik
	var isMarried bool
	fmt.Println(isMarried) //çalışınca false başlangıç sonucunu verir
	/*	fmt.Println(name)
		fmt.Println(age)
		fmt.Println(isMarried)
		fmt.Println(weight)
		fmt.Println(height)
	*/
}
