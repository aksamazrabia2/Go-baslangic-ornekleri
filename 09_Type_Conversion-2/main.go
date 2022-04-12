package main

import "fmt"

func main() {
	/*var x int8 = 10
	var y int16 = 10
	fmt.Println(x + y)  Hata verir türleri farklı*/

	/*x := 10
	y := 10.0

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)

	fmt.Println(float64(x) + y)*/
	var x int8 = 127
	//var y int16
	//y = x Bu şekilde hata verir.Type conversion yapmam lazım
	var y int16
	y = int16(x) //type(value)
	fmt.Println(y)
	/*x := 10
	y := "10"
	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)
	fmt.Println(x + int(y))
	type conversion yöntemiyle herhangi bir int değeri string değerine dönüşmez.
	*/
	/*num1 := 106
	str1 := string(num1)
	fmt.Printf("%v %T\n", num1, num1)
	fmt.Println()
	fmt.Printf("%v %T\n", str1, str1)
*/
}
