package main

import "fmt"

func main() {
	x := 100
	y := 20
	z := 30

	fmt.Printf("%b %d %o", x, y, z)

	name, age := "Rabia", 26
	fmt.Print("Benim adım ", name, ", ve ben ", age, " yaşındayım. ")
	fmt.Println("Benim adım", name, "ve ben", age, "yaşındayım. ")
	fmt.Printf("Benim adım %v ve ben %v yaşındayım. ", name, age)

}
