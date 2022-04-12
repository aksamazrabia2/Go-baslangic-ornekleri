package main

import "fmt"

func main() {

	//if <boolean expression> {code} x%2==0(false)

	/*	x := 27

		if x%2 == 0 {
			fmt.Println(x, "çift sayıdır")

		}
		if x%2 != 0 {
			fmt.Println(x, "tek sayıdır")

		}
	*/

	/*if !true {//false olursa mesaj gösterilir.
		fmt.Println("Mesaj Gosterilecek")
	}*/
	/*if 5 > 3 {
		fmt.Println("Mesaj gösterilecek mi?")
	}*/
	/*x := 27

	if x%2 == 0 {
		fmt.Println(x, "çift sayıdır")

	} else {
		fmt.Println(x, "tek sayıdır")
	}*/
	/*if true {
		fmt.Println("Mesaj Gösterilecek ")
	}*/
	x := -5

	if x < 0 {
		fmt.Println(x, "negatif sayıdır")

	} else if x%2 == 0 {
		fmt.Println(x, "çift sayıdır")
	} else {
		fmt.Println(x, "tek sayıdır")
	}
	//if <boolean expression> {code} else if <boolean expression> else {code}
}
