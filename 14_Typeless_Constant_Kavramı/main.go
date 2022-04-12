package main

import "fmt"

func main() {

	/*const x int16 = 5
	fmt.Printf("%T, %v", x, x) */

	/*const x int8 = 3
	var y int16 = 12
	fmt.Printf("%T, %v", x, x)
	fmt.Printf("%T, %v", y, y)*/
	//fmt.Printf("%T, %v", x+y, x+y) hata verir

	/*const x = 3
	var y int16 = 12
	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", y, y)
	fmt.Printf("%T, %v\n", x+y, x+y)//int16(x)*/

	/*const x = int16(5.2 + 4.8)
	fmt.Printf("%T, %v", x, x)*/

	/*const x int32 = 3
	const y float32 = 5.6
	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", y, y)
	//fmt.Printf("%T, %v\n", x+y, x+y) //Hata verir.*/

	/*const x = 3
	const y = 5.6
	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", y, y)
	fmt.Printf("%T, %v\n", x+y, x+y) //Hata vermez*/

	const x = 3
	const y = 5.6
	const z = true
	const t = "test"
	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", y, y)
	fmt.Printf("%T, %v\n", z, z)
	fmt.Printf("%T, %v\n", t, t)

}
