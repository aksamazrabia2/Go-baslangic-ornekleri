package main

import "fmt"

func main() {
	/*
		fmt.Print("Merhaba")
		fmt.Println("Merhaba")
	*/
	/*Yukardaki iki satır sırayla çalısır sorunsuz bir şekilde ancak yan yana yazar.aşagıdaki gibi yerlerini değiştirisek*/
	/*ln den dolayı bir alt satıra iner ve kodu alt satırdan yazar.*/
	/*
		fmt.Println("Merhaba")
		fmt.Print("Merhaba")
		fmt.Printf("Merhaba")
	*/
	name := "Rabia"
	/*
		fmt.Print(name)
		fmt.Println(name)
		fmt.Printf(name)
	*/
	fmt.Print("Benim adım", name)
	fmt.Println("")
	fmt.Println("Benim adım", name)
	fmt.Printf("Benim adım %v %T", name, name)
	fmt.Println("")
	fmt.Printf("Benim adım %X %T", name, name)
}
