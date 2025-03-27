package main

import "fmt"

// ! for => is only construct for looping in GOLANG

func main() {

	// implementing while loop using for bcz there is While keyword in GOLANG

	i:=1

	for (i<=5){
		fmt.Println("helloWorld")
		i++
	}

	//! sometimes we use while loop for infinite things that syntax is easy 
	// for {
	// 		fmt.Println("hello")
	// }

	// ! classic for loop 

	for i:=0;i<=5;i++{
		fmt.Println(i)
	}

	// ! continue  is used for skipping the current iteration 
	// ! break is used for breaking the loop 


	for i:=0;i<=10;i++{
		if i==5{
			continue
		}
		fmt.Println(2*i)
	}

	// ! there is range keyword can be used for from 0 to whatever number you specify that will be only excluded 

	for i:= range 10 {
		fmt.Println("hello World",i)

	}
	
}