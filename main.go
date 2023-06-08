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
