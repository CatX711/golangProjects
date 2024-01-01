package main

import "fmt"
import "strings"
import "time"

func main(){
	var mynumber float32 = 3.14159268
	var mynumber2 float32 = 237123.13231233
	var yesorno string

	fmt.Println("\n\nHello, this is my first program using my modified vim!")
	addnums(mynumber, mynumber2)
	fmt.Printf("Do you like vim? (yes/no) ")
	fmt.Scan(&yesorno)

	if strings.ToLower(yesorno) == "yes"{
		fmt.Println("Great!!!")
	} else if strings.ToLower(yesorno) == "no"{
		time.Sleep(2 * time.Second)
		fmt.Printf("\x1b[2J")
		time.Sleep(3 * time.Second)
		fmt.Println("Get out.")
	} else{
		fmt.Println("Huh?")
	}	
}

func addnums(num1, num2 float32){
	sum := num1 + num2
	fmt.Println(sum)
}
