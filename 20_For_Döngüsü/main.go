package main

import "fmt"

func main() {
	/*	for i := 1; i <= 10; i++ {
			fmt.Println(i)
		}
	*/

	//for <init statement>; <condition>; <post statement>{..........}
	/*i := 0
	for ; i <= 10; i += 5 {
		fmt.Println(i)
	}
	fmt.Println(i)
	*/
	/*	for {//Infiniti Loop
			fmt.Println("Benim adım Rabia")
		}
	*/
	/*for i := 0; true; i += 5 {
		fmt.Println(i)
	}*/

	/*i := 10
	for i >= 0 {
		fmt.Println(i)
		i--
	}*/
	// continue-break deyimleri

	/*for i := 0; i <= 10; i++ {
		if i%3 == 0 {
			continue// bu koşulu sağladığın zaman döngünün başına git
		}
		fmt.Println(i)
	}*/
	for i := 0; i <= 10; i++ {
		if i == 3 {
			break //döngüden çık
		}
		fmt.Println(i)
	}

}
