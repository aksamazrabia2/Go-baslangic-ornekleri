package main

import "fmt"

func main() {
	x := 10
	y := 10.0

	fmt.Printf("%v %T \n", x, x)
	fmt.Printf("%v %T\n", y, y)
	//fmt.Println(x + y) veri tipler ifarklı diye hata  verir.

	//Type conversion type(value)=>int(y)=10.0=>10
	fmt.Println(x + int(y)) //değikenin aslı değişmez

	fmt.Printf("%v %T\n", y, y)

}
