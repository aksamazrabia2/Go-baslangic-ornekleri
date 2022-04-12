//1)İki rakam arasında toplama,çıkarma ve çarpma işleminin yapıldığı bir fonksiyon yazınız. package main
/*package main

import "fmt"

func main() {

	x, y := 10, 4
	sum, dif, prod := calculation(x, y)
	fmt.Println("Toplam:", sum)
	fmt.Println("Fark:", dif)
	fmt.Println("Çarpma:", prod)

}
func calculation(num1, num2 int) (int, int, int) {
	sum := num1 + num2
	dif := num1 - num2
	prod := num1 * num2

	return sum, dif, prod
}
*/
//2)Kullanıcı tarafından girilen nota göre geçtiniz veya kaldınız geri dönüşünü yazdırınız.
/*package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("Lütfen aldığınız notu giriniz:")
	grade, _ := getGrade()

	var result string

	if grade > 50 {
		result = "Geçtin"
	} else {
		result = "Kaldın"
	}
	fmt.Println(result)

}
func getGrade() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	input = strings.TrimSpace(input)
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
	}
	return num, nil

}*/

//3) 1 ile yüz arasında bir sayıyı tahmin etme ugulaması yazınız.Toplma tahmin hakkınız 10 olsun.
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	target := numRand(1, 100)

	fmt.Println("1 ile 100 arasındaki sayıyı bulmay çalışın")

	reader := bufio.NewReader(os.Stdin)

	for attempts := 0; attempts < 10; attempts++ {

		fmt.Println(10-attempts, "hakkınız kaldı")
		fmt.Println("Lütfen tahmininizi yapınız:")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
		}
		if num > target {
			fmt.Println("Tahmininiz daha büyük lütfen daha küçük bir sayı giriniz")
		} else if num < target {

			fmt.Println("Tahmininiz daha küçük lütfen daha büyük bir sayı giriniz")
		} else {
			fmt.Println("Doğru tahmin", attempts, " seferde buldunuz")
			break
		}
	}

}
func numRand(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
