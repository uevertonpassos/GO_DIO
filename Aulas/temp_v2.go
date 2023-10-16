package main

import "fmt"

const ebulicaoF = 212.0

func main(){
	tempF := ebulicaoF
	tempC := (tempF - 32)*5/9 

	fmt.Printf("A ebulição da água acontecem em %.2f °F e em %.2f em °C", tempF,
	tempC)


}