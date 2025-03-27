package main

import "fmt" // fmt is inbuilt package in Go which has function called Println

func main(){
	fmt.Println("Hello World")
	fmt.Print("GO hello")
}

//  if you want to build the exe file => go build 01_helloWorld/main.go
// an executable file will be created with name as => main.exe
//  and to execute the file  ./main.go 

// to directly run without building it => go run 01_helloWorld/main.go  