package main

import "fmt"

//***Fonksiyon isimlendirme***
//ilk karakter harf olmalı
//camel case-->mySum,myBestFunction
//paket dışı-->ilk harf büyük
func main() {

	/*var x, y, sum int
	x = 5
	y = 10
	sum = x + y
	fmt.Printf("%d ve %d toplamı %d\n", x, y, sum)

	x = 7
	y = 11
	sum = x + y
	fmt.Printf("%d ve %d toplamı %d\n", x, y, sum)*/
	//fonsiyon ile işlemlerimizi daha modüler hale getiririz.(moduler programming)
	//kodlar daha kolay bir şekilde okunur(readeble code)
	//comlex işlemi basitleştirebiliriz(from complex to simple)
	/*
		fmt.Println(sum(5, 10))
		merhaba()
		fmt.Println(sum(3, 7))
		fmt.Println(sum(2, 5))
		fmt.Println(sum(50, 10))

		merhaba()
		merhaba()
	*/
	//return vs print
	z := sum(5, 10)
	fmt.Println(z)
	sum2(6, 11)

	merhaba2("rabia", 26)

}

//func funcName(params) return type{ code}
func sum(x, y int) int {

	return x + y

}
func sum2(x, y int) {
	fmt.Println(x + y)
}

func merhaba() {
	fmt.Println("Benim adım rabia")
}

func merhaba2(name string, age int) {
	fmt.Printf("Adım %s, yaşım %d", name, age)
}
