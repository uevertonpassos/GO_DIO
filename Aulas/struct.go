package main

import "fmt"

type pessoa struct{
	nome string
	idade int

}

func main(){
	fmt.Println(pessoa{"Uev", 27})

}