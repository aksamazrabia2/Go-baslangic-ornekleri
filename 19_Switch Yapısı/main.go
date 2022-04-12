package main

import "fmt"

func main() {

	//switch
	grade := 4

	switch grade {
	case 5: //if grade==5 {fmt.Println("Pekiyi")}
		fmt.Println("Pekiyi")

	case 4:
		fmt.Println("İyi")
		y := 100
		fmt.Println(y)
	case 3:
		fmt.Println("Orta")

	case 2:
		fmt.Println("Geçer")

	case 1:
		fmt.Println("Başarısız")

	default:
		fmt.Println("Geçersiz Not")
	}

	/*if grade == 5 {
		fmt.Println("Pekiyi")
	} else if grade == 4 {
		fmt.Println("İyi")
	} else if grade == 3 {
		fmt.Println("Orta")
	} else if grade == 2 {
		fmt.Println("Geçer")
	} else if grade == 1 {
		fmt.Println("Başarısız")
	} else {
		fmt.Println("Geçersiz Not")
	}*/
	//fmt.Println(grade)
}
