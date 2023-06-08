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
package main

import "fmt"

//int float bool qiymatga o'zgartirish
func main()  {
	var a float64
	fmt.Scan(&a)

	fmt.Printf("Int qiymati=%d\n",int(a))
	fmt.Printf("Float qiymat=%f\n",a)
	var x bool
	if a!=0 {
		x=true
	}
	fmt.Printf("bool qiymat=%t\n",x)

}
package main

import "fmt"

//bahoni tekshiruvchi kod
func main()  {
	var a int
	fmt.Scan(&a)
	if a>80 && a<100 {
		fmt.Println("Bahoyingiz 5",)
	} else if a>60 && a<80 {
		fmt.Println("Bahoyingiz 4")
	} else if a>40 && a<60 {
		fmt.Println("Bahoyingiz 3")
	} else if a>0 && a<40 {
		fmt.Println("Bahoyingiz 2")
	} else {
		fmt.Println("0 dan kichkina son kiritish mumkin emas")
	}
} 
package main 

import "fmt"

//harfmi yoki harf emasmi kodi
func main() {
	var n rune
	fmt.Scanf("%c", &n)

	if (n >= 'A' && n <= 'Z') || (n >= 'a' && n <= 'z') {
		fmt.Println("harf")
	} else {
		fmt.Println("harf emas")
	}
}