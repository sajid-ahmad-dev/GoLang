package main

import (
	"fmt"
	"time"
)

func main() {

	// switch is used when we have too many conditions to check 
	// for one codition we use  if else 

	i:=4

	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("other")

	}

	// ! mulitple conditions 

	switch time.Now().Weekday() {
	case time.Saturday,time.Sunday: 
		fmt.Println("its weekend")
	default:
		fmt.Println("its workday")
	}

	// ! type switch 

	         // whoamI := func(i interface{}) {
		                                // whoamI is a variable that is assigned an anonymous function.
		
		                               // := is the short declaration operator, which means Go infers the type of whoamI based on the function assigned to it.
		
		                               // func(i interface{}) { defines an anonymous function (a function without a name):
		
		                              // It takes a single parameter i.
		
		                             // The type of i is interface{}, which is Go's empty interface.
		
    whoamI:=func(i interface {}){// The empty interface (interface{}) can hold a value of any type (similar to any in other languages like TypeScript or void* in C).
		switch t:=i.(type){
		case int:
			fmt.Println("its an integer")
		case string:
			fmt.Println("its an string")
		case bool:
			fmt.Println("its an bool")
		default:
			fmt.Println("other ",t)
		}

		

	}
	whoamI(994.6)



}