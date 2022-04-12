package main

import "fmt"

func main() {
	/*x := 10

	if x := -5; x < 0 {
		fmt.Println(x, "negatif sayıdır.")
	} else if x%2 == 0 {
		fmt.Println(x, "çift sayıdır.")

	} else {
		fmt.Println(x, "tek sayıdır.")
	}

	fmt.Println(x)*/
	x := 25
	if x < 0 {
		fmt.Println(x, "negatif sayıdır.")
	} else {
		if x%2 == 0 {
			fmt.Println(x, "çift sayıdır.")
		} else {
			fmt.Println(x, "tek sayıdır.")
		}
	}

}
