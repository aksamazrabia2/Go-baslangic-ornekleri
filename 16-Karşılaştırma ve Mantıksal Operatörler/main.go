//Comparison operators
//== Equal !=Not equal
//< Less than > greater than
//<= Less than or equal >=Greater than equal

//logical Operators
//&& conditional AND p && q is "if p then q else false"
//|| conditional OR p || q is "if p then true else q"
// ! NOT !p is "not p"

package main

import "fmt"

func main() {
	/*	x, y := 3, "b" //3,7
		fmt.Printf("%T, %v\n", x == y, x == y)
		fmt.Printf("%T, %v\n", x != y, x != y)
		fmt.Printf("%T, %v\n", x < y, x < y)
		fmt.Printf("%T, %v\n", x > y, x > y)
		fmt.Printf("%T, %v\n", x >= y, x >= y)
		fmt.Printf("%T, %v\n", x <= y, x <= y)*/

	x, y := 15, 20
	set1 := (x == y) //false
	set2 := (x > y)  //false
	/*fmt.Printf("%T, %v\n", set1, set1)
	fmt.Printf("%T, %v\n", set2, set2)*/
	set3 := true

	fmt.Printf("%v\n", (set1 && set2)) //AND--> sadece 2 durum true-->true
	fmt.Printf("%v\n", (set1 || set2)) //OR --> sadece 2 durum false-->false
	fmt.Printf("%v\n", (set3 || set2))
	fmt.Printf("%v\n", (!set3))
}
