package main

import "fmt"

func main(){

	for i := 0;i<=10;i++{ //prestar atenção na declaração
		if i > 5 || i< 8{
			fmt.Println(i, "LANA")
		}else{
			fmt.Println("aqui é:", i)
		}
	}
}
