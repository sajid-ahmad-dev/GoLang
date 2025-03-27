package main

import "fmt"

const name string = "Macbook"

const anotherName = "M4"

// employed:=false  //! not acceptable

func main() {

	//! also if you have declared the constant variable , you also need to assign a value at the time of declaration 

	const age = 85 // this is constant value named as age , you cannot reassign it

	//  you can also declare it outside the main function , but the only exception is that you can not use this syntax below

	//! age:=85  outside the main function this syntax is not valid

	fmt.Println(age)
	fmt.Println(name)


	//!  you can also do the constant grouping , like for some value you know that this will remain constant throught the whole process

	const (
		port int =5500
		name="webserver" // if you do not give the type it infers 
		


	)
	fmt.Println(name,port)

}