//addition +=>15+10=>15,10 operand,+ operator
//substruction -
//product *
//division /
//remainder %
package main

import "fmt"

func main() {
	/*x, y := 15, 10
	fmt.Printf("%T ,%v\n", x, x)
	fmt.Printf("%T ,%v\n", y, y)
	fmt.Printf("Toplama:%T,%v\n", (x + y), (x + y))
	fmt.Printf("Çıkarma:%T,%v\n", (x - y), (x - y))
	fmt.Printf("Çarpma:%T,%v\n", (x * y), (x * y))
	fmt.Printf("Bölme:%T,%v\n", (x / y), (x / y))
	z := 5.0 / 2
	fmt.Printf("%T ,%v\n", z, z)
	fmt.Printf("Mod:%T,%v\n", (x % y), (x % y))*/
	//Increment ++,Decrement --
	x := 10

	fmt.Println(x)
	x = x + 1
	fmt.Println(x)
	x++
	fmt.Println(x)
	x = x - 1
	fmt.Println(x)
	x--
	fmt.Println(x)
}
