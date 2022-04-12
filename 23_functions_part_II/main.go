/*package main

import "fmt"

func main() {

	merhaba("Rabia", 26) //argument function run
}

func merhaba(name string, age int) { //prametre function write

	fmt.Printf("Adım %s,yaşım %d", name, age)

}*/
/*package main

import "fmt"

func main() {
	fmt.Println(result(60))
}
func result(grade int) string {
	if grade >= 50 {
		//fmt.Println("IF KOŞULU İÇİNDEYİM")
		return "geçtiniz"

	}

	return "kaldınız"

}*/

/*package main

import "fmt"

func main() {
	merhaba("rabia", 26)

	name := "meryem"
	age := 22
	fmt.Printf("Adım %s, yaşım %d", name, age)//buradaki name ve age in fonksiyon içindekilerle alakası yok

}
func merhaba(name string, age int) { //prametre function write

	fmt.Printf("Adım %s,yaşım %d\n", name, age)

}*/
/*package main

import (
	"fmt"
	"time"
)

func main() {

	var x int = 10

	fmt.Println(x)

	var moment time.Time = time.Now()//Now--> method

	fmt.Println(moment)

}*/
/*package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Print("Lütfen Sınav Sonucunuzu Giriniz:")
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n') //_blank identifier
	fmt.Println(value)
}*/

package main

import "fmt"

//104/5===>20-4 verecek
func main() {

	bölüm, kalan := bölme(104, 5)
	fmt.Println(bölüm, kalan)
}
func bölme(bölünen, bölen int) (bölüm, kalan int) {
	bölüm = bölünen / bölen
	kalan = bölünen % bölen

	return bölüm, kalan

}
