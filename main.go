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
package main

import "fmt"
// Raqamlar yig'indisini toping
func main() {
	var n int
	var sum int
	fmt.Scan(&n)
	for i := 0; i <= n; i++ {
		sum += i
	}
	fmt.Printf("Raqamlar yig'indisi: %d\n", sum)
}
package main

import "fmt"
//sonlani o'rnini almashtirish dasturi
 func main() {
     var number int
     fmt.Print("2 xonali son kiriting: ")
     fmt.Scan(&number)
     n := number % 10
     new := (number / 10) % 10
     newNum := n*10 + new
     fmt.Printf(" %d\n",newNum )
 }
 package main

 import "fmt"
 func main() {
    var a int
    fmt.Println("7 xonali son kiriting: ")
    fmt.Scan(&a)

    num1:= a % 10000000
    num2 := (a / 1000000) % 10
	num3 := (a / 100000) % 10
	num4 := (a / 10000) % 10
	num5 := (a / 1000) % 10
	num6 := (a / 100) % 10
	num7 := (a / 10) % 10
    newNumber := firstDigit*10 + secondDigit

    fmt.Printf("%d\n",newNumber)
}
package main

import "fmt"
//Sekundni chiqaruvchi dastur
func main() {
    var seconds int
    fmt.Print("sekundlarni kiriting: ")
    fmt.Scan(&seconds)

    hours := seconds / 3600
    minutes := (seconds % 3600) / 60
    seconds = seconds % 60
   fmt.Printf("%d hour %d minute %d second\n", hours, minutes, seconds)
}
package main

import "fmt"
//karra kara jadvali dastur
func main() {
	for n:=1; n<=10; n++ {
		for m:=1; m<=10; m++ {
			fmt.Printf("%v*%v=%v\t",m,n,m*n)
		}
		fmt.Println("  ")
	}
}
package main

import "fmt"
//userdan berilgan songacha karra kara jadvalini chiqarish
func main() {
	var a int
	fmt.Scan(&a)
	for n:=1; n<=a; n++ {
		for m:=1; m<=a; m++ {
			fmt.Printf("%v*%v=%v\t",m,n,m*n)
		}
		fmt.Println("  ")
	}
}
package main

import "fmt"
//harflarni boshi va oxirini o'zgartiruvchi dastur
func main() {
	var input string
	fmt.Scan(&input)

	res := string(input[len(input)-1]) + string(input[1:len(input)-1]) + string(input[0])
	fmt.Println(res)
}
package main

import "fmt"
// raqamlar yig'indisni topiish
func main() {
	var number, res int
	fmt.Scan(&number)
	for number != 0 {
		last := number % 10
		res += last
		number = (number - last) / 10
		fmt.Println(res)
	}

}
	//2.variant
	func main() {
		var number, res int
		fmt.Scan(&number)
		res = number / 100000
		res += int(number/10000) % 10
		res += int(number/1000) % 10
		res += int(number/100) % 10
		res += int(number/10) % 10
		res += number % 10
		fmt.Println(res)

}
package main

import "fmt"
// ikkilanganini chiqaradigan dastur
	func main() {
		var a, b, c int
		fmt.Scan(&a, &b, &c)
		if a < b && b < c {
			fmt.Println(2*a, a*b, 2*c)
		} else {
			fmt.Println(-a, -b, -c)
		}
	}
package main

import "fmt"
//string (nodir) o'rtasiga chiqaruvchi dastur

	func main() {
		var string1, string2 string
		fmt.Scan(&string1, &string2)
		res := string1[:len(string1)/2] + string2 + string1[len(string1)/2:]
		fmt.Println(res)
}
package main

import "fmt"
//1 dan n gacha bolgan juft sonlarni sonini chiqarish dasturi

	func main() {
		var n int
		fmt.Scan(&n)
		count := 0
		for i := 1; i <= n; i++ {
			if i%2 == 0 {
				count++
			}
		}
		fmt.Printf("Juft=%v\n", count)
}
package main

import "fmt"
// kiritilgan string uzunligi 4 ga bo'linsa ikkinchi yarmini teskari qilib o'zgartirish
	func main() {
		var str, res string
		fmt.Scan(&str)
		if len(str)%4 == 0 {
			for i := len(str) / 2; i < len(str); i++ {
				res = string(str[i]) + res
			}
		}
		fmt.Println(str[:len(str)/2] + res)
	}
package main 

import "fmt"
//so'zni oxirgi va bosh harfni kattalashtirish
func main() {
	var str string
	fmt.Scan(&str)

	str = strings.ToUpper(string(str[0])) + strings.ToUpper(string(str[len(str)-1]))
	fmt.Println(str)
}
package main 

import "fmt"
// so'zda nechta a borligini chiqaruvchi dastur
	func main() {
		var str string
		fmt.Scan(&str)
		count := 0
		for _, ch := range str {
			if ch == 'a' {
				count++
			}
		}
		fmt.Println(count)
	}
package main 

import "fmt"
//so'zlarni indekslari toq bolsa katta harf bilan chiqaruvchi dastur
func main() {
	var input string
	fmt.Print("So'zni kiriting: ")
	fmt.Scan(&input)

		for i := 0; i < len(input); i++ {
			if i%2 == 0 {
				fmt.Printf("%c", input[i]-32)
			} else {
				fmt.Printf("%c", input[i])
			}
		}
	}
package main
import "fmt"
//Foydalanuvchi bitta gap va n son kiritadi. Siz esa, n chi belgini olib tashlang va shu gapni ekranga chiqaring.
func main() {
	var str string
	var n int
	fmt.Scan(&str, &n)
	if n > 0 && n <= len(str) {
		fmt.Printf("%d-chi belgi: %c\n", n, str[n-1])
	}
}
package main

import "fmt"
// Foydalanuvchi kiritgan string teskari tartibda ekranga chiqaring
	func main() {
		var str string
		fmt.Scan(&str)
		for i := len(str) - 1; i >= 0; i-- {
			fmt.Print(string(str[i]))
		}
	}
package main

import "fmt"
// palindrome
func main() {
	matn := "kiyik"
	fmt.Println(palindrome(matn))
}
func palindrome(str string) bool {
	for i, j := 0, len(str)-1; i < len(str)/2; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
			break
		}
	}
	return true
}