package main

import "fmt"

func main() {
	age := 0

	if age >= 18 {
		fmt.Println("You are an adult")
	} else if age > 12 {
		fmt.Println("You are a teenager")
	} else {
		fmt.Println("You are a kid")
	}

	var user="admin"
	user="Rega"
	var hasPermission=true

	if user=="admin" && hasPermission{
		fmt.Println("You are an admin , therefore permission to enter")
	} else {
		fmt.Println("Regular User , not allowed")
	}

	//! we can also declare the variable here in   if-else statement

	if salary:=60000; salary>=50000 {
		fmt.Println("Dream job")
	} else if salary >40000{
		fmt.Println("Median Salary")
	}else {
		fmt.Println("guzara layak mehnat kr pasina baha ")
	}

	//! GOLANG does not have ternary operator till to this date => 28 march 2025



}