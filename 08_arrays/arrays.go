package main

import "fmt"

func main() {

	// how to declare the arrays 
	//An array in Go has a fixed size, and its length must be specified at declaration:

	var nums [4]int
	var str [2]string
	var val [4]bool
	fmt.Println(val)
	str[0]="anna"
	fmt.Println(str[1])
	fmt.Println(str)
	fmt.Println(nums)
	fmt.Println(len(val))

	// ! to declare and array within in one line

	values := [4]int {0,1,2,3}

	fmt.Println(values)

	// ! 2d Array

	TwodValues:= [2][2]int {{1,2},{4,5}}   // declares 2 rows and 2 columns 
	fmt.Println(TwodValues)

	another2d:= [3][4]int {{1,25,5},{7,8}}
	fmt.Println(another2d)
	

	// fixed size , that is predictable 
	// memory optimization 
	// constant time access ( element accessed in constant time )

	// but if you do not know the size , for dynamic memory allocation , you can use slices constructor

}