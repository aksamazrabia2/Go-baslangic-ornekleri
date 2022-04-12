package main

import (
	"fmt"
)

func main() {
	/*r := 3.0
	const pi float64 = 3.14
	areaOfCircle := pi * (math.Pow(r, 2)) //burda 3.14 kendi başına sabit
	fmt.Println(areaOfCircle)*/

	/*const x int = 2
	const y float32 = 3.14
	const z string = "test"
	const t bool = false
	fmt.Printf("%T %v\n", x, x)
	fmt.Printf("%T %v\n", y, y)
	fmt.Printf("%T %v\n", z, z)
	fmt.Printf("%T %v\n", t, t)*/

	/*	const x int = 2
			//x=5 hata alırız.const'lar değişmez.
	 		fmt.Printf("%T %v\n", x, x)*/

	//const--->compile time, var--->run time

	/*var x = math.Pow(3, 4)
	//const x = math.Pow(3, 4) böyle yazsak hata alırız.
	fmt.Printf("%T %v\n", x, x)*/

	y := 3
	x := 3 //yerine const x=y dersek hata verir.
	fmt.Printf("%T %v\n", y, y)
	fmt.Printf("%T %v\n", x, x)

}
