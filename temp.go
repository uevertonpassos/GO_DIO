package main

import "fmt"

const ebulicaoF  = 212.0

func main(){
	var tempF  = ebulicaoF
	var tempC  = (tempF - 32)*5/9 

	fmt.Println("A temperatura de ebulição da água em °F = ",tempF)
	fmt.Println("A temperatura de ebulição da água em °C = ",tempC)


}