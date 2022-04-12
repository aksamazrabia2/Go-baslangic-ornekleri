//1-) studentName --> John Doe, grade-->77 ,isPassed-->true değişkenlerini 3 farklı yöntemle oluşturup,çıktısını alınız.

//--------------1.Soru-----------------

/*package main

import "fmt"

func main() {

	/*1.Yöntem
	var studentName string = "John Doe"
	var grade int = 77
	var isPassed bool = true

	fmt.Println(studentName)
	fmt.Println(grade)
	fmt.Println(isPassed)
*/

/*2.Yöntem
var studentName = "John Doe"
var grade = 77
var isPassed = true

fmt.Println(studentName)
fmt.Println(grade)
fmt.Println(isPassed)
*/

/*3.Yöntem
	studentName := "John Doe"
	grade := 77
	isPassed := true

	fmt.Println(studentName)
	fmt.Println(grade)
	fmt.Println(isPassed)

}*/

//-----------------2.Soru------------------

package main

import "fmt"

func main() {

	//var studentName, grade, isPassed = "John Due", 77, true
	studentName, grade, isPassed := "John Due", 77, true
	fmt.Println(studentName)
	fmt.Println(grade)
	fmt.Println(isPassed)

}

//2-Yukarıda belirtilen değişkenleri tek satır içerisinde tanımlayınız.
//3-"Declaration","Assign","Initialization","Initial Value" kavramlarını açıklayınız.(Terminoloji)
//4"Statically Typed" ve "Dynamically Typed" ifadelerini Go ve Python üzerinden gösteriniz.
//5-":=" vs "=" aradaki farkı gösteriniz.
