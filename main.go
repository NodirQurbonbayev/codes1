package main

import "fmt"

//Juft yoki toqligini aniqlaydigan
func main() {
	var a int
	fmt.Println("Sonni kiriting:")
	fmt.Scan(&a)
	if a%2 == 1 {
		fmt.Println("Toq=", a)
	} else {
		fmt.Println("Juft=", a)
	}

}
package main

import "fmt"

// uzunligini aniqlaydigan
func main() {
	var str string
	fmt.Scan(&str)
	length := len(str)
	fmt.Println(length)
}
package main 

import "fmt"
//Probelni o'chirib string uzunligini chiqaradi 
func main() {
	var str1, str2 string
	fmt.Println("Stringni kiriting:")
	fmt.Scan(&str1, &str2)

	strlength, err := fmt.Println(str1 + str2)
	if err != nil {
		fmt.Println("xato bor", err)
	}
	fmt.Printf("uzunligi:%v\n", strlength)
}
package main 

import "fmt"
 //ikkita sonni bo'lganda yuqoriga qarab yaxtlishlash
func main() {
    var a, b float64
    fmt.Println("birinchi sonni kiriting: ")
    fmt.Scan(&a)
    fmt.Println("ikkinchi sonni kiriting: ")
    fmt.Scan(&b)

    if b != 0 {
        natija := a / b
        fmt.Printf("Natija: %.2f\n", natija)
    } else {
        fmt.Println("Xatolik: nolga bo'lish mumkin emas!")
    }
}