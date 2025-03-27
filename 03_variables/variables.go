package main

import "fmt"

func main (){

	var name string="Golang"   // if declared a variable then you need to use it , it mandatory , if you do no use it then remove it  
	
	// OR you can also declare the variables like this 

	var name2="Golang2"   // here this name2 infer that this is string bcz the value we gave is in " " likewise you can see the bool value too

	var condition=true

	var age int =30

	// this above syntax is used in cases where you do not want to assign a value , but want to declare the variable 

	var employed bool

	var price float32



	fmt.Println(age)

	fmt.Println(condition)
	
	fmt.Println(name2)
	
	fmt.Println(name)

	// there is also a ShortHand Syntax for variables to declare

	myname:="Sajid"  // it will also infer the variable type

	hours:=6

	employed=false

	price=85369.55
	fmt.Println(price)
	fmt.Println(myname)
	fmt.Println(employed)
	fmt.Println(hours)
}