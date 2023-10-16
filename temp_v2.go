package main

import "fmt"

const ebulicaoF  = 212.0

func main(){
	tempF  := ebulicaoF
	tempC  := (tempF - 32)*5/9 

	fmt.Printf("A ebulição da água acontecem em %g °F e em %g em °C", tempF,
	tempC)


}