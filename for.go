package main

import "fmt"

func main(){

	//for e IF em go
	for i := 1; i<=10; i++{

		if i%2 == 0{
			fmt.Println("par")
		}else {
			fmt.Println("ímpar!")
		}
	}
}